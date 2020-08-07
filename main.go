package main

import (
  "github.com/synbioz/go_api/config"
  "github.com/synbioz/go_api/models"
  "log"
  "net/http"
)

func main() {
  config.DatabaseInit()
  router := InitializeRouter()

  // Populate database
  models.NewUser(&models.User{email: "jeremy.beaucousin@test.fr", firstname: "Jérémy", lastname: "BEAUCOUSIN"})

  log.Fatal(http.ListenAndServe(":8080", router))
}