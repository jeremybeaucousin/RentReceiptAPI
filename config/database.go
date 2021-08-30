package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	var databaseUrl string
	databaseUrl = os.Getenv("DATABASE_URL")
	if len(databaseUrl) == 0 {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		databaseUrl = os.Getenv("DATABASE_URL")
	}

	db, err = sql.Open("postgres", databaseUrl+"?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createUsersTable()
}

func createUsersTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id serial, email text not null, firstname varchar(20) not null, lastname varchar(20) not null, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
