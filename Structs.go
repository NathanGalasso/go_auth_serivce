package Users

import (
	"errors"
	"fmt"
	"log"
)

type userData struct {
	uName string
	uPass string
}

func NewUser(uName string, uPass string) (*userData, error) {
	// If no name was given, return an error with a message.

	if uName == "" {
		return nil, errors.New("Empty Name")
	}

	// If no password was given, return an error with a message.

	if uPass == "" {
		return nil, errors.New("Empty Name")
	}

	p := userData{uName, uPass}
	return &p, nil
}

func main() {
	NathanUser, err := NewUser("Nathan", "password")
	if err != nil {
		log.Fatal(err)
	}
	// if no error was returned print the username and password
	fmt.Println("Username:", NathanUser.uName, "\nPassword:", NathanUser.uPass)
}
