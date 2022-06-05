package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitPGConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("[PqConfig] Error loading .env file.")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatal("[PqConfig] Error getting postgres connection.", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("[PqConfig] Error pinging database.", err)
	}

	return db
}
