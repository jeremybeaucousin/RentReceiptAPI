package controllers

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jeremybeaucousin/RentReceiptAPI/models"

	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const propertyIdKey string = "propertyId"

func PropertiesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	ownerId := getQueryVar(r, ownerIdKey)

	json.NewEncoder(w).Encode(models.AllProperties(ownerId))
}

func PropertiesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	ownerId := getQueryVar(r, ownerIdKey)

	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var property models.Property

	error = json.Unmarshal(body, &property)

	if error != nil {
		log.Fatal(error)
	}

	savedProperty := models.NewProperty(ownerId, &property)

	json.NewEncoder(w).Encode(savedProperty)
}

func PropertiesShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	ownerId := getQueryVar(r, ownerIdKey)

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}

	property := models.FindPropertyById(ownerId, id)

	if property == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(property)
	}

}

func PropertiesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	ownerId := getQueryVar(r, ownerIdKey)

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	property := models.FindPropertyById(ownerId, id)
	if property == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		body, error := ioutil.ReadAll(r.Body)
		error = json.Unmarshal(body, &property)
		models.UpdateProperty(property)
		json.NewEncoder(w).Encode(property)
		w.WriteHeader(http.StatusOK)
		if error != nil {
			log.Fatal(error)
		}
	}
	if error != nil {
		log.Fatal(error)
	}
}

func PropertiesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	ownerId := getQueryVar(r, ownerIdKey)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}
	property := models.FindPropertyById(ownerId, id)

	if property == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		models.DeletePropertyById(property)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(property)
	}
}
