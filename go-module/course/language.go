package course

import (
	"fmt"
)

func Apply() string {
	fmt.Println("Please enter your first name")
	var firstName string
	if _, err := fmt.Scan(&firstName); err != nil {
		return ""
	}

	fmt.Println("Please enter your last name")
	var lastName string
	if _, err := fmt.Scan(&lastName); err != nil {
		return ""
	}

	fmt.Printf("Hi %s %s thank you for joinging code 4 easy \n", firstName, lastName)

	return firstName + " " + lastName
}
