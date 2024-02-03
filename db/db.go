package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func GetDB() *sql.DB {
	envError := godotenv.Load((".env"))
	if envError != nil {
		fmt.Printf("Could not load env file")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", os.Getenv("DB_STRING_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
