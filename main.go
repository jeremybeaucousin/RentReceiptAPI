package main

import (
	"log"
	"net/http"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
