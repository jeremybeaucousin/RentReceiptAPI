package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDb *gorm.DB

func GormDatabaseInit() {
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
	gormDb, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

// Getter for db var
func GormDb() *gorm.DB {
	return gormDb
}
