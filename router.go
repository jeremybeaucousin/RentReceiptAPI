package main

import (
	"github.com/gorilla/mux"
	"github.com/jeremybeaucousin/RentReceiptAPI/controllers"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /users/ to /users
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/owners").HandlerFunc(controllers.OwnersIndex)
	router.Methods("POST").Path("/owners").HandlerFunc(controllers.OwnersCreate)
	router.Methods("GET").Path("/owners/{id}").HandlerFunc(controllers.OwnersShow)
	router.Methods("PUT").Path("/owners/{id}").HandlerFunc(controllers.OwnersUpdate)
	router.Methods("DELETE").Path("/owners/{id}").HandlerFunc(controllers.OwnersDelete)

	router.Methods("GET").Path("/owners/{ownerId}/properties").HandlerFunc(controllers.PropertiesIndex)
	router.Methods("POST").Path("/owners/{ownerId}/properties").HandlerFunc(controllers.PropertiesCreate)
	router.Methods("GET").Path("/owners/{ownerId}/properties/{id}").HandlerFunc(controllers.PropertiesShow)
	router.Methods("PUT").Path("/owners/{ownerId}/properties/{id}").HandlerFunc(controllers.PropertiesUpdate)
	router.Methods("DELETE").Path("/owners/{ownerId}/properties/{id}").HandlerFunc(controllers.PropertiesDelete)

	router.Methods("GET").Path("/owners/{ownerId}/properties/{propertyId}/tenants").HandlerFunc(controllers.TenantsIndex)
	router.Methods("POST").Path("/owners/{ownerId}/properties/{propertyId}/tenants").HandlerFunc(controllers.TenantsCreate)
	router.Methods("GET").Path("/owners/{ownerId}/properties/{propertyId}/tenants/{id}").HandlerFunc(controllers.TenantsShow)
	router.Methods("PUT").Path("/owners/{ownerId}/properties/{propertyId}/tenants/{id}").HandlerFunc(controllers.TenantsUpdate)
	router.Methods("DELETE").Path("/owners/{ownerId}/properties/{propertyId}/tenants/{id}").HandlerFunc(controllers.TenantsDelete)

	return router
}
