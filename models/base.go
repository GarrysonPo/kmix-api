package models

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	db = conn
}

func CloseDB() {
	db.Close()
}

func GetDB() *sql.DB {
	return db
}
