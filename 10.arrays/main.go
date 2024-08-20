package main

import "fmt"

func main() {
	// 🔥 array
	var number [10]int
	number[0] = 10
	fmt.Println(number)
	fmt.Println(number[0])
	fmt.Println(len(number)) // len() return the length of the array

	var name [20]string
	fmt.Println(name)
	fmt.Println(name[0])

	var values [5]bool
	fmt.Println(values)
	fmt.Println(values[0])

	// 🔥 declare & initilize at same time
	marks := [3]float32{23.4, 344.53, 2.32}
	fmt.Println(marks)

	// 🔥 2D array
	// var grade [2][2]int
	grade := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(grade)
	fmt.Println(grade[0])
	fmt.Println(grade[1])
	fmt.Println(grade[1][0])
}
