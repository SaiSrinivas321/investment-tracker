package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"investment-tracker/config"
)

var db *sql.DB

func InitDB() error {
	var err error
	dbHost := config.GetEnv("DB_HOST")
	dbPort := config.GetEnv("DB_PORT")
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbName := config.GetEnv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connected successfully!")
	return nil
}

func GetDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
