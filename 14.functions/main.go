package main

import "fmt"

func func1() {
	fmt.Println("Takes nothing, return nothing")
}

func func2() string {
	fmt.Println("Takes nothing, return something")
	return "Called func2"
}

func func3(x int, y int) {
	fmt.Println("Takes something, return nothing")
}
func func4(x, y int) int { // x & y both are int
	fmt.Println("Takes something, return something")
	return x + y
}
func main() {
	// ------ function --------
	x := 1
	y := 2

	func1() // 🔥 takes nothing, return nothing

	func2() // 🔥 takes nothing, return something

	func3(x, y) // 🔥 takes something, return nothing

	func4(x, y) // 🔥 takes something, return something
}
