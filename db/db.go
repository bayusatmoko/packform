package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
	envError := godotenv.Load((".env"))
	if envError != nil {
		fmt.Printf("Could not load env file")
		os.Exit(1)
	}

	var err error

	db, err = sql.Open("postgres", os.Getenv("DB_STRING_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
