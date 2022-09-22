package databases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (dbc *DBConfig) ConnectPostgres(ctx context.Context) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, dbc.genConnStr())
	if err != nil {
		return nil, fmt.Errorf("(dbc *DBConfig)ConnectPostgres #1 \n Error:%s \n", err.Error())
	}

	if err = db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("(dbc *DBConfig)ConnectPostgres #2 \n Error:%s \n", err.Error())
	}

	return db, nil
}

func (dbc *DBConfig) ConnectClickHous() (*sql.DB, error) {
	db, err := sql.Open("clickhouse", dbc.genConnStr())
	if err != nil {
		return nil, fmt.Errorf("(dbc *DBConfig)ConnectClickHous #1 \n Error:%s \n", err.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("(dbc *DBConfig)ConnectClickHous #2 \n Error:%s \n", err.Error())
	}

	return db, nil
}

func (dbc *DBConfig) genConnStr() string {
	switch dbc.Driver {
	case Postgres:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbc.User,
			dbc.Password,
			dbc.Host,
			dbc.Port,
			dbc.DBName,
			dbc.SSLmode)
	case ClickHouse:
		return fmt.Sprintf(
			"tcp://%s:%s?username=%s&password=%s&database=%s&debug=%s",
			dbc.Host,
			dbc.Port,
			dbc.User,
			dbc.Password,
			dbc.Name,
			dbc.Debug,
		)
	default:
		return ""
	}
}
