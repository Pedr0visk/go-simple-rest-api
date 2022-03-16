package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connInfo := connection{
		DBHost: os.Getenv("POSTGRES_URL"),
		DBPort: os.Getenv("POSTGRES_PORT"),
		DBUser: os.Getenv("POSTGRES_USER"),
		DBPass: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
	}

	db, err = sql.Open("postgres", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB is open\n")
	}

	// check if we can ping our DB
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")
	}
}

// Take our connection struct and convert to a string for our db connection info
func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.DBHost, info.DBPort, info.DBUser, info.DBPass, info.DBName)

}
