package itemService

import (
	"database/sql"
)

type clickHouseDB struct {
	conn *sql.DB
}

func NewClickHouseDB(conn *sql.DB) *clickHouseDB {
	return &clickHouseDB{conn: conn}
}

func (pdb *clickHouseDB) Insert(logs []string) error {

	return nil
}
