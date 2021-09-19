package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getQueryVar(r *http.Request, key string) int {
	vars := mux.Vars(r)
	ownerId, error := strconv.Atoi(vars[key])

	if error != nil {
		log.Fatal(error)
	}
	return ownerId
}
