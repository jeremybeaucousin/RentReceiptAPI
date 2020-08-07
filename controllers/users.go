package controllers

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/synbioz/go_api/models"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  json.NewEncoder(w).Encode(models.AllUsers())
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var user models.User

	error = json.Unmarshal(body, &user)

	if error != nil {
		log.Fatal(error)
	}

	models.NewUser(&user)

	json.NewEncoder(w).Encode(user)
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}

	user := models.FindUserById(id)

	json.NewEncoder(w).Encode(user)
}

func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}

	body, err := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	user := models.FindUserById(id)

	error = json.Unmarshal(body, &user)

	models.UpdateUser(user)

	json.NewEncoder(w).Encode(user)
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		log.Fatal(error)
	}

	error = models.DeleteUserById(id)
}