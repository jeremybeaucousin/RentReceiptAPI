package config

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
)

var db *sql.DB

func DatabaseInit() {
  var err error

  db, err = sql.Open("postgres", "user=rentReceiptGenerator password=Jonathan dbname=rentReceiptGenerator")

  if err != nil {
    log.Fatal(err)
  }

  // Create Table cars if not exists
  createUsersTable()
}

func createUsersTable() {
  _, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id serial, email text not null, firstname varchar(20) not null, lastname varchar(20) not null, constraint pk primary key(id))")

  if err != nil {
    log.Fatal(err)
  }
}

// Getter for db var
func Db() *sql.DB {
  return db
}