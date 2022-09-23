package databases

import (
	"context"
	"fmt"

	//"github.com/ClickHouse/clickhouse-go/v2"
	ch "github.com/leprosus/golang-clickhouse"
	//"github.com/ClickHouse/clickhouse-go/v2"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/jackc/pgx/v4"
)

func (dbc *DBConfig) ConnectPostgres(ctx context.Context) (*pgx.Conn, error) {
	//db, err := pgx.Connect(ctx, dbc.genConnStr())
	//"172.17.0.1" c ним из контенера сервера цепляется в базе
	db, err := pgx.Connect(ctx, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.DBName))

	if err != nil {
		return nil, fmt.Errorf("(dbc *DBConfig)ConnectPostgres #1 \n Error:%s \n", err.Error())
	}

	if err = db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("(dbc *DBConfig)ConnectPostgres #2 \n Error:%s \n", err.Error())
	}

	return db, nil
}

func (dbc *DBConfig) ConnectClickHous() *ch.Conn {
	//db, err := clickhouse.Open(&clickhouse.Options{
	//	Addr: []string{"127.0.0.1:8123"},
	//	Auth: clickhouse.Auth{
	//		Database: "default",
	//		Username: "default",
	//		Password: "password",
	//	},
	//})
	db := ch.New("127.0.0.1", 8123, "default", "")

	//
	//if err != nil {
	//	return nil, fmt.Errorf("(dbc *DBConfig)ConnectClickHous #1 \n Error:%s \n", err.Error())
	//}

	//if err = db.Ping(context.Background()); err != nil {
	//	return nil, fmt.Errorf("(dbc *DBConfig)ConnectClickHous #2 \n Error:%s \n", err.Error())
	//}

	return db
}

func (dbc *DBConfig) genConnStr() string {
	switch dbc.Driver {
	case Postgres:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbc.User,
			dbc.Password,
			dbc.Host,
			dbc.Port)
	case ClickHouse:
		return fmt.Sprintf(
			"%s://%s:%s?username=%s&password=%s&database=%s&debug=%s",
			dbc.Driver,
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
