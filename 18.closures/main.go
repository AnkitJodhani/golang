package main

import "fmt"

/* In the closure
- outer function variable will remain in the memory after execution of inner function
*/
func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}
