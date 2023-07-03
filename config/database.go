package config

import (
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDb *gorm.DB

func GormDatabaseInit() {
	var err error
	var databaseUrl string

	if len(os.Getenv("INSTANCE_HOST")) == 0 {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		databaseUrl = os.Getenv("DATABASE_URL")
	} else {
		var (
			dbUser    = os.Getenv("DB_USER")
			dbPwd     = os.Getenv("DB_PASS")
			dbName    = os.Getenv("DB_NAME")
			dbPort    = os.Getenv("DB_PORT")
			dbTCPHost = os.Getenv("INSTANCE_HOST")
		)
		databaseUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
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
