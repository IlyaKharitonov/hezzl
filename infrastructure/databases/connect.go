package databases

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func(db *Database)Connect()error{
	var err error
	ctx := context.Background()
	db.DB, err = pgx.Connect(ctx,db.genConnStr())

	if err != nil{
		return fmt.Errorf("db.(db *Database)Connect #1 \n Error:%s \n", err.Error())
	}

	if err = db.DB.Ping(ctx); err!= nil {
		return fmt.Errorf("db.(db *Database)Connect #2 \n Error:%s \n", err.Error())
	}

	return nil
}

func (db *Database)genConnStr()string{
	switch db.Driver {
	case Postgres:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			db.User,
			db.Password,
			db.Host,
			db.Port,
			db.DBName,
			db.SSLmode)
	case ClickHouse:
		return fmt.Sprintf(
			"tcp://%s:%s?username=%s&password=%s&database=%s&debug=%s",
			db.Host,
			db.Port,
			db.User,
			db.Password,
			db.Name,
			db.Debug,
		)
	default:
		return ""
	}
}

