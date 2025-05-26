package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_PORT"), 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_PASSWORD"), 
	os.Getenv("DB_NAME"),
)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	fmt.Println("✅Connected to database")
	DB = db
}
