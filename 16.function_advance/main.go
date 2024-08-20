/*
🔸 operate function:
    - This function takes two integers (a and b) and a function (operation) that itself takes two integers
	and returns an integer. This allows you to pass different operations like addition, multiplication, etc.,
	as functions.

🔸 multiplier function:
	- This function returns another function. The returned function multiplies its input by a factor that
	 was specified when the multiplier function was first called.


🔸 accumulator function:
	- This is an example of a closure. The accumulator function returns a function that adds its input to
	a running total. The running total (sum) is maintained in the environment of the returned function,
	which "remembers" the value of sum across multiple calls.

🔸 Anonymous function:
	- The example shows how you can define and immediately invoke an anonymous function in Go.
	This is often useful for quick, inline operations.

*/

package main

import (
	"fmt"
)

// 1. Function that takes a function as a parameter
func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 2. Function that returns another function
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 3. Closure example: a function that captures an external variable
func accumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Example of passing a function as a parameter
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	fmt.Println("Add 5 + 3:", operate(5, 3, add))           // Add 5 + 3: 8
	fmt.Println("Multiply 5 * 3:", operate(5, 3, multiply)) // Multiply 5 * 3: 15

	// Example of returning a function from another function
	timesTwo := multiplier(2)
	timesThree := multiplier(3)

	fmt.Println("4 * 2:", timesTwo(4))   // 4 * 2: 8
	fmt.Println("4 * 3:", timesThree(4)) // 4 * 3: 12

	// Example of a closure
	acc := accumulator()

	fmt.Println("Add 1 to accumulator:", acc(1)) // Add 1 to accumulator: 1
	fmt.Println("Add 5 to accumulator:", acc(5)) // Add 5 to accumulator: 6
	fmt.Println("Add 2 to accumulator:", acc(2)) // Add 2 to accumulator: 8

	// Anonymous function used directly
	result := func(x, y int) int { return x - y }(10, 4)
	fmt.Println("10 - 4:", result) // 10 - 4: 6
}
