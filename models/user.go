package models

import (
	"database/sql"
	"log"

	"github.com/jeremybeaucousin/RentReceiptAPI/config"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func NewUser(user *User) *User {
	if user == nil {
		log.Fatal(user)
	}

	error := config.Db().QueryRow("INSERT INTO users (email, firstname, lastname) VALUES ($1,$2,$3) RETURNING id;", user.Email, user.Firstname, user.Lastname).Scan(&user.Id)

	if error != nil {
		log.Fatal(error)
	}

	return user
}

func FindUserById(id int) *User {
	var user User

	row := config.Db().QueryRow("SELECT * FROM users WHERE id = $1;", id)

	error := row.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname)

	switch {
	case error == sql.ErrNoRows:
		log.Printf("no user with id %d", id)
		return nil
	case error != nil:
		log.Fatal(error)
		return nil
	default:
		log.Printf("user is %s\n", user.Firstname)
		return &user
	}
}

func AllUsers() *[]User {
	var users []User

	rows, error := config.Db().Query("SELECT * FROM users")

	if error != nil {
		log.Fatal(error)
	}

	for rows.Next() {
		var user User

		error := rows.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname)

		if error != nil {
			log.Fatal(error)
		}

		users = append(users, user)
	}

	// Close rows after all readed
	defer rows.Close()

	return &users
}

func UpdateUser(user *User) {
	stmt, error := config.Db().Prepare("UPDATE users SET email=$1, firstname=$2, lastname=$3 WHERE id=$4;")

	_, error = stmt.Exec(user.Email, user.Firstname, user.Lastname, user.Id)

	if error != nil {
		log.Fatal(error)
	}
}

func DeleteUserById(id int) {
	stmt, error := config.Db().Prepare("DELETE FROM users WHERE id=$1;")

	if error != nil {
		log.Fatal(error)
	}

	_, error = stmt.Exec(id)

	if error != nil {
		log.Fatal(error)
	}
}
