package main

import "fmt"

func main() {

	// 🔥 iterating slice
	num := []int{1, 2, 3, 4, 5}
	// "i" has index & "j" has value
	for i, j := range num {
		fmt.Println(i, j)
	}

	// 🔥 iterating maps
	dic := map[string]string{"name": "Ankit", "fname": "hasubhai"}
	for key, value := range dic {
		fmt.Println(key, value)
	}

	// 🔥 iterating string
	name := "Ankit Jodhania"
	for i := range name {
		fmt.Println(name[i]) // print ASCII  value
	}

}
