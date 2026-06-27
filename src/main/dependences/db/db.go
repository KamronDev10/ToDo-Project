package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	err := godotenv.Load("src/resources/.dev.env")
	if err != nil {
		log.Fatal("env yuklanmadi:", err)
	}

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB ochilmadi:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB ulanmadi:", err)
	}

	log.Println("DB ga ulandi ✅")
	return db, nil
}
