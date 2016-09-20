package databaseService

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db MySQLDriver;

func init() {
	db, err := sql.Open("mysql", os.GetEnv("MYSQL_CONNECTIONG_STRING"))
	
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}
}