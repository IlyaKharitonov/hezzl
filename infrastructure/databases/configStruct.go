package databases

import (
	"github.com/jackc/pgx/v4"
)

//for pg and ch
type Database struct {
	DB *pgx.Conn
	Driver string `json:"driver"`
	Host string 	`json:"host"`
	Port string  `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Name string `json:"name"`
	DBName string `json:"dbName"`
	SSLmode string `json:"sslMode"`
	Debug string `json:"debug"`
}






