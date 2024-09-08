package main

import "fmt"

func fact(n int) int {

	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)

}

func main() {
	var num int
	fmt.Println("Please enter the number: ")
	fmt.Scanln(&num)
	factorial := fact(num)
	fmt.Printf("Factorial of %v is: %v", num, factorial)
}
