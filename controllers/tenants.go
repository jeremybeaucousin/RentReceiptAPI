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

func TenantsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllTenants())
}

func TenantsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var tenant models.Tenant

	error = json.Unmarshal(body, &tenant)

	if error != nil {
		log.Fatal(error)
	}

	savedTenant := models.NewTenant(&tenant)

	json.NewEncoder(w).Encode(savedTenant)
}

func TenantsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}

	tenant := models.FindTenantById(id)

	if tenant == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tenant)
	}

}

func TenantsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	tenant := models.FindTenantById(id)
	if tenant == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		body, error := ioutil.ReadAll(r.Body)
		error = json.Unmarshal(body, &tenant)
		models.UpdateTenant(tenant)
		json.NewEncoder(w).Encode(tenant)
		w.WriteHeader(http.StatusOK)
		if error != nil {
			log.Fatal(error)
		}
	}
	if error != nil {
		log.Fatal(error)
	}
}

func TenantsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}
	tenant := models.FindTenantById(id)

	if tenant == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		models.DeleteTenantById(tenant)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tenant)
	}
}
