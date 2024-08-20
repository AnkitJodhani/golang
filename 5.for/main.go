package main

import "fmt"

func main() {
	// 🔥 for loop
	fmt.Println("for loop")

	count := 10
	for i := 0; i < count; i++ {
		if i == 3 {
			// break
			continue
		}
		fmt.Println(i)
	}

	// 🔥 infinite loop
	// for {
	// 	fmt.Println("keep printing")
	// }

	//  🔥 range - go 1.22
	fmt.Println("--- range ----")
	for i := range 3 {
		fmt.Println(i)
	}
}
