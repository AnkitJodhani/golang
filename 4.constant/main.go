package main

import "fmt"

const age = 90

var name string = "Ankit"

// name:= "Ankit" //  "short variable declaration"  syntax not allowed in global space

func main() {

	// 🔥 constant
	const price float64 = 343.2
	fmt.Println(price)

	// 🔥 constants groping - to create multiple constant
	const (
		portNumber = 3000
		cpu        = 4
		hostname   = "localhost"
	)
	fmt.Println(portNumber, cpu, hostname)
}
