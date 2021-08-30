package main

import (
	"github.com/gorilla/mux"
	"github.com/jeremybeaucousin/RentReceiptAPI/controllers"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /users/ to /users
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/users").HandlerFunc(controllers.UsersIndex)
	router.Methods("POST").Path("/users").HandlerFunc(controllers.UsersCreate)
	router.Methods("GET").Path("/users/{id}").HandlerFunc(controllers.UsersShow)
	router.Methods("PUT").Path("/users/{id}").HandlerFunc(controllers.UsersUpdate)
	router.Methods("DELETE").Path("/users/{id}").HandlerFunc(controllers.UsersDelete)

	return router
}
