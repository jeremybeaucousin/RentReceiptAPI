package config

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDb *gorm.DB

func GormDatabaseInit() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if len(os.Getenv("INSTANCE_CONNECTION_NAME")) == 0 {
		databaseUrl := os.Getenv("DATABASE_URL")
		gormDb, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
		if err != nil {
			panic("failed to connect database" + databaseUrl)
		}
	} else {
		var (
			dbUser                 = os.Getenv("DB_USER")
			dbPwd                  = os.Getenv("DB_PASS")
			dbName                 = os.Getenv("DB_NAME")
			instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
			usePrivate             = os.Getenv("PRIVATE_IP")
		)
		dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=disable", instanceConnectionName, dbUser, dbPwd, dbName)
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Error to parse config")
		}
		var opts []cloudsqlconn.Option
		if usePrivate != "" {
			opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
		}
		d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
		if err != nil {
			log.Fatalf("Error get connector")
		}
		// Use the Cloud SQL connector to handle connecting to the instance.
		// This approach does *NOT* require the Cloud SQL proxy.
		config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName)
		}
		databaseUrl := stdlib.RegisterConnConfig(config)
		gormDb, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN:        dsn,
		}))
		if err != nil {
			panic("failed to connect database" + databaseUrl)
		}
	}
}

// Getter for db var
func GormDb() *gorm.DB {
	return gormDb
}
