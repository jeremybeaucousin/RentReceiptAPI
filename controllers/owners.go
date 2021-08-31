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

func OwnersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllOwners())
}

func OwnersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var owner models.Owner

	error = json.Unmarshal(body, &owner)

	if error != nil {
		log.Fatal(error)
	}

	savedOwner := models.NewOwner(&owner)

	json.NewEncoder(w).Encode(savedOwner)
}

func OwnersShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}

	owner := models.FindOwnerById(id)

	if owner == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(owner)
	}

}

func OwnersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	owner := models.FindOwnerById(id)
	if owner == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		body, error := ioutil.ReadAll(r.Body)
		error = json.Unmarshal(body, &owner)
		log.Print(owner)
		models.UpdateOwner(owner)
		json.NewEncoder(w).Encode(owner)
		w.WriteHeader(http.StatusOK)
		if error != nil {
			log.Fatal(error)
		}
	}
	if error != nil {
		log.Fatal(error)
	}
}

func OwnersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}
	owner := models.FindOwnerById(id)

	if owner == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		models.DeleteOwnerById(owner)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(owner)
	}
}
