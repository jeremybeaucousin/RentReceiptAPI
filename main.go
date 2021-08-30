package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()
	var port string
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
