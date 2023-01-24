package constract

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type DBConnect struct {
	DBType     string
	DBUser     string
	DBIp       string
	DBPort     string
	DBDatabase string
}

func (receiver DBConnect) MySQLConnect(user string, ip string, port string, database string) *sqlx.DB {
	dbConnection, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println(err)
	}
	return dbConnection
}
