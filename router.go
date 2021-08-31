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

	router.Methods("GET").Path("/owners").HandlerFunc(controllers.OwnersIndex)
	router.Methods("POST").Path("/owners").HandlerFunc(controllers.OwnersCreate)
	router.Methods("GET").Path("/owners/{id}").HandlerFunc(controllers.OwnersShow)
	router.Methods("PUT").Path("/owners/{id}").HandlerFunc(controllers.OwnersUpdate)
	router.Methods("DELETE").Path("/owners/{id}").HandlerFunc(controllers.OwnersDelete)

	return router
}
