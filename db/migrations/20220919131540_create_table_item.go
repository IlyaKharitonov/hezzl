
package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Up is executed when this migration is applied
func Up_20220919131540(txn *sql.Tx) {
	_,err := txn.Exec(`CREATE TABLE IF NOT EXISTS items
(
    id serial 		PRIMARY KEY,
    campaign_id 	INTEGER NOT NULL REFERENCES campaigns(id),
    name 			VARCHAR(100) NOT NULL,
    description	 	VARCHAR(100),
    priority 		INTEGER DEFAULT(1),
    removed boolean DEFAULT(FALSE),
    created_at 		TIMESTAMP DEFAULT(CURRENT_TIMESTAMP)
);

CREATE INDEX id_idx ON items(id);
CREATE INDEX campaign_id_idx ON items(campaign_id);
CREATE INDEX name_idx ON items USING HASH (name);`)
	if err != nil {
		log.Printf("Up_20220919131540: %s\n", err)
	}
	fmt.Println("Up_20220919131540 OK")

}

// Down is executed when this migration is rolled back
func Down_20220919131540(txn *sql.Tx) {
	_,err := txn.Exec(`DROP TABLE items;`)
	if err != nil {
		log.Printf("Down_20220919131540: %s\n", err)
	}
	fmt.Println("Down_20220919131540 OK")
}
