package databaseService

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hacksoc-manchester/www/helpers/validator"
	"log"
	"os"
)

type userEntry struct {
	FirstName            string
	LastName             string
	Email                string
	SubscribedToArticles bool
	SubscribedToEvents   bool
}

var db *sql.DB

func init() {

	if os.Getenv("MYSQL_CONNECTION_STRING") == "" {
		log.Println("Environment variable MYSQL_CONNECTION_STRING is not assigned.")
		return
	}

	db, _ = sql.Open("mysql", os.Getenv("MYSQL_CONNECTION_STRING"))

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
}

// CreateUser adds the specified entry to the User table.
func CreateUser(firstName, lastName, email string, subscribedToArticles, subscribedToEvents bool) error {
	if !validator.IsValidName(firstName) {
		return errors.New(`Name "` + firstName + `" is not valid.`)
	}

	if !validator.IsValidName(lastName) {
		return errors.New(`Name "` + lastName + `" is not valid.`)
	}

	if !validator.IsValidEmail(email) {
		return errors.New(`Email "` + email + `" is not valid.`)
	}

	stmt, _ := db.Prepare(`insert into User(FirstName, LastName, Email, SubscribedToArticles, SubscribedToEvents)
		values (?, ?, ?, ?, ?)`)

	if _, err := stmt.Exec(firstName, lastName, email, subscribedToArticles, subscribedToEvents); err != nil {
		return errors.New(`Email "` + email + `" is already subscribed.`)
	}

	return nil
}

// GetUser retrieves the user with the specified email.
func GetUser(email string) (*userEntry, error) {
	stmt, _ := db.Prepare("select * from User where Email = ?")
	user := new(userEntry)
	err := stmt.QueryRow(email).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.SubscribedToArticles,
		&user.SubscribedToEvents)

	if err != nil {
		return nil, errors.New(`Could not find user with email "` + email + `".`)
	}

	return user, nil
}

// ExistsUser determines whether a user with the specified email exists.
func ExistsUser(email string) bool {
	user, _ := GetUser(email)

	return user != nil
}

// DeleteUser removes the specified entry from the User table.
func DeleteUser(email string) error {
	stmt, _ := db.Prepare("delete from User where Email = ?")

	if _, err := stmt.Exec(email); err != nil {
		return errors.New(`Could not delete user with email "` + email + `".`)
	}

	return nil
}
