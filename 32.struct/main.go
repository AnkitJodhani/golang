package main

import (
	"fmt"

	"ankit.com/project/user"
)

func main() {
	var firstName string
	var lastName string
	var emailId string
	fmt.Println("Enter your first name")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your email id")
	fmt.Scanln(&emailId)

	ankit, err := user.New(firstName, lastName, emailId)
	// ✅ validation
	if err != nil {
		panic(err)
	}
	fmt.Println(ankit)
	fmt.Println(*ankit)

}
