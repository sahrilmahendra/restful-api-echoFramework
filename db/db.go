package db

import (
	"database/sql"
	"echo-rest/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	/* template connection :
	username:password@protocol(address:port)/dbname */
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("tidak dapat terkoneksi ke database.. koneksi error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func createCon() *sql.DB {
	return db
}
