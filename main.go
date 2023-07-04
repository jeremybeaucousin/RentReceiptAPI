package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
	"github.com/jeremybeaucousin/RentReceiptAPI/models"
	"github.com/rs/cors"
)

func main() {
	config.GormDatabaseInit()
	config.GormDb().AutoMigrate(&models.Owner{}, &models.Tenant{}, &models.Property{})

	router := InitializeRouter()

	var port string
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	log.Printf("Used port is %s", port)
	allowedOrigins := strings.Split(os.Getenv("ORIGIN_RENT_REICEIPT_GENERATOR"), ",")
	log.Printf("Origins active is %s", allowedOrigins)

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
