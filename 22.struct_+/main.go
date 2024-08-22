package main

import "fmt"

func main() {
	// annonymos -  don't have to give the name- just like object in js
	house := struct {
		name     string
		number   int
		area     float32
		landmark string
	}{"Shyam socity", 18, 720.00, "Krypton school"}

	fmt.Println(house)
	fmt.Println(house.area)
	fmt.Println(house.landmark)

}
