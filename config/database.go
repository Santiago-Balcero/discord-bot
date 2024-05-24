package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	db, err := sql.Open(DbDriver, DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
