package main

import (
	"fmt"
)

// sum is variadic function that can take n number of parameter
func sum(nums ...int) int { // "...int" will collect each sperate value & create slice
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func test(slice ...interface{}) { // "interface{}" can accept any kind(type) of value
	for _, value := range slice {
		fmt.Println(value)
	}
}

func main() {
	// 🔥 Sending multiple value
	fmt.Println(sum(2, 4, 23, 1))

	// 🔥 Sending multiple value - kind of "Spread operator"
	list2 := []int{2, 4, 23, 1} // slice
	fmt.Println(sum(list2...))  // sending each value as parameter - "Spreading"

	// 🔥 Sending multiple value of different type
	test("Ankit", 19, 202.3, 8.07, false)
}

/*
Note: 💦 spreading &  ☔ collecting

💦 spreading
something... ➡️ It means you are "spreading" slice or array

☔ collecting
...something ➡️ It means you are "collecting" & creating slice

*/
