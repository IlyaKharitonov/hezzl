
package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Up is executed when this migration is applied
func Up_20220919131517(txn *sql.Tx) {
	_,err := txn.Exec(`
CREATE TABLE IF NOT EXISTS campaigns
(
	id serial 	PRIMARY KEY,
	name 		VARCHAR(100) NOT NULL UNIQUE
);
CREATE INDEX idx_id ON campaigns(id);
INSERT INTO campaigns (name) VALUES ('First campaign');
`)
	if err != nil {
		log.Printf("Up_20220919131517: %s\n", err)
	}
	fmt.Println("Up_20220919131517 OK")
}

// Down is executed when this migration is rolled back
func Down_20220919131517(txn *sql.Tx) {
	_,err := txn.Exec(`DROP TABLE campaigns CASCADE;`)
	if err != nil {
		log.Printf("Down_20220919131517: %s\n", err)
	}
	fmt.Println("Down_20220919131517 OK")
}
