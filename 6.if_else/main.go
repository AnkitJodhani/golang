package main

import "fmt"

func main() {
	// 🔥 simple if else statment
	var age int

	fmt.Println("Please enter your age: ")
	fmt.Scan(&age)

	if age > 18 {
		fmt.Println(" You can drive the car/bike 🚗 ")
	} else if age == 18 {
		fmt.Println(" Please hold on for a year to drive 🚗 ")
	} else {
		fmt.Println(" Please have a patience 🙏 ")
	}

	// 🔥 simple if else statment
	var role string
	var hasPermission bool

	fmt.Println("Who are you? ")
	fmt.Scan(&role)
	fmt.Println("Do you have permission? ")
	fmt.Scan(&hasPermission)

	if role == "admin" && hasPermission == true {
		fmt.Println("🙌 Verification successful")
	} else {
		fmt.Println("😔 Sorry but we can't let you go in")
	}

	//  Note: Go does not have tutrnay operator yet..

}
