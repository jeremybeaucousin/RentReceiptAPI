package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"github.com/jeremybeaucousin/RentReceiptAPI/models"
	"github.com/rs/cors"
)

func main() {
	config.DatabaseInit()
	config.GormDatabaseInit()
	config.GormDb().AutoMigrate(&models.Owner{}, &models.Tenant{}, &models.Property{})

	router := InitializeRouter()

	var port string
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Printf("Used port is %s", port)
	log.Printf("Origin active is %s", os.Getenv("ORIGIN_RENT_REICEIPT_GENERATOR"))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ORIGIN_RENT_REICEIPT_GENERATOR")},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
