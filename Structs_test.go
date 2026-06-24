package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestNewUser(t *testing.T) {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")   //Set this system environment variable to your SQL username
	cfg.Passwd = os.Getenv("DBPASS") //Set this system environment variable to your SQL password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "authproject" //When creating the database for the project make the name "authproject" to match or change to your database name

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	//Ensure that the database can be accesses and is connected
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	//Dummy values for new user to be created
	userName := "Nathan"
	userPassword := "Password"
	user, err := NewUser(userName, userPassword)
	if err != nil {
		t.Errorf("New userName = %q, New userPassword = %q, Error = %v", userName, userPassword, err)
	}
	//Print the successfully created new user
	fmt.Printf("New user name: %s,New user pass: %s, New user id: %d\n", user.uName, user.uPass, user.userid)
}
