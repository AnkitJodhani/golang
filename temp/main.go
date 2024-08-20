package main

import "fmt"

func accumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	a := accumulator()
	fmt.Println(a(1))
	fmt.Println(a(2))
}
