package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type userData struct {
	userid int64
	uName  string
	uPass  string
}

// Function that is used to add a new user to the users table.
// The function will return the id of the new user entry that was created, returning 0 for an error.
func NewUser(uName string, uPass string) (userData, error) {
	//A userdata variable to hold the newly created user
	var user userData
	// If no name was given, return an error with a message.
	if uName == "" {
		return user, errors.New("Empty Name")
	}
	// If no password was given, return an error with a message.
	if uPass == "" {
		return user, errors.New("Empty Name")
	}
	//Insert the userdata values into the database table
	result, err := db.Exec("INSERT INTO users (username, pass) VALUES (?, ?)", uName, uPass)
	if err != nil {
		return user, fmt.Errorf("NewUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("NewUser: %v", err)
	}
	//Get the details about the newly created user from the database table to ensure it was created successfully.
	row := db.QueryRow("select * from users where id = ?", id)
	if err := row.Scan(&user.userid, &user.uName, &user.uPass); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("NewUser: %d: no such User", id)
		}
		return user, fmt.Errorf("NewUser %d: %v", id, err)
	}
	if err != nil {
		return user, fmt.Errorf("NewUser: %v", err)
	}
	//Return the successfully created userData
	return user, nil
}

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "authproject"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	//Ping the database to make sure it was connected successfully
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// Dummy values for new user to be created.
	NathanUser, err := NewUser("Nathan", "password")
	if err != nil {
		log.Fatal(err)
	}
	// if no error was returned print the username and password
	fmt.Println("Username:", NathanUser.uName, "\nPassword:", NathanUser.uPass, "\nNew User ID:", NathanUser.userid)
}
