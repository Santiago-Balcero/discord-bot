package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost,
		DbPort,
		DbUser,
		DbPassword,
		DbName,
	)
	db, err := sql.Open(DbDriver, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
