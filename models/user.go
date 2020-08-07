package models

import (
  "github.com/synbioz/go_api/config"
  "log"
  "time"
)

type User struct {
  Id			int       `json:"id"`
  Email			string    `json:"email"`
  Firstname		string    `json:"firstname"`
  Lastname		string    `json:"lastname"`
}

type Users []User

func NewUser(user *User) {
	if user == nil {
		log.Fatal(user)
	}
		
	error := config.Db().QueryRow("INSERT INTO users (email, firstname, lastname) VALUES ($1,$2,$3) RETURNING id;", user.Email, user.Firstname, user.Lastname).Scan(&c.Id)

	if err != nil {
		log.Fatal(error)
	}
}

func FindUserById(id int) *User {
	var user User

	row := config.Db().QueryRow("SELECT * FROM users WHERE id = $1;", id)
	error := row.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname)

	if error != nil {
		log.Fatal(error)
	}

	return &user
}

func AllUsers() *Users {
	var users Users

	rows, error := config.Db().Query("SELECT * FROM users")

	if error != nil {
		log.Fatal(error)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var user Uar

		error := rows.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname)

		if error != nil {
			log.Fatal(error)
		}

		users = append(users, c)
	}

	return &users
}

func UpdateUser(user *User) {
	stmt, error := config.Db().Prepare("UPDATE users SET email=$1, firstname=$2, lastname=$3 WHERE id=$4;")

	if error != nil {
		log.Fatal(error)
	}

	_, error = stmt.Exec(user.Email, user.Firstname, user.Lastname, user.Id)

	if error != nil {
		log.Fatal(error)
	}
}

func DeleteUserById(id int) error {
	stmt, error := config.Db().Prepare("DELETE FROM users WHERE id=$1;")

	if error != nil {
		log.Fatal(error)
	}

	_, error = stmt.Exec(error)

	return error
}