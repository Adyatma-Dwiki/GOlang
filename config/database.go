package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// connect to db
	db, err := sql.Open("mysql", "root:root@/go_product")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
