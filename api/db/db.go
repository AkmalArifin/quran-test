package db

import (
	"database/sql"
	"os"
	"time"
)

var DB *sql.DB

func InitDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	var err error

	dataSource := user + ":" + password + "@tcp(localhost:3306)/srq_project"
	DB, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err.Error())
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
