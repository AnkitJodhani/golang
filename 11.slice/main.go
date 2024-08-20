package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice -> dynamic array

	// 🔥 uninitilized slice is nil
	var student []int
	fmt.Println(student == nil)
	fmt.Println(len(student))

	// 🔥 declare & initilize slice
	var teacher = []int{1, 123234, 3, 3}
	fmt.Println(teacher)

	// 🔥 declare, initilize with ZERO
	var member = make([]int, 2)
	fmt.Println(member)
	fmt.Println(member == nil)

	// 🔥 declare, initilize with ZERO & Initial CAPACITY
	// 🗒️ Note: once initial capacity get full & then slice automatically double the capacity
	var sports = make([]int, 2, 5)
	fmt.Println(sports)
	fmt.Println(cap(sports)) // capacity
	sports = append(sports, 2)
	sports = append(sports, 3)
	sports = append(sports, 4)
	// sports = append(sports, 5)
	// sports = append(sports, 6)
	fmt.Println(sports)
	fmt.Println(cap(sports)) // capacity

	// -------------------------------------why make() function? --------------------------------------
	/*
		Even though slices are dynamic, specifying the length and capacity when using make() allows you to:
		1) Control the initial state of the slice.
		2) Optimize performance by reducing the need for reallocations.
		3) Preallocate memory when you know the approximate size the slice will grow to.

	*/
	computer := make([]int, 3, 5)
	computer = append(computer, 2)
	computer = append(computer, 3)
	fmt.Println(computer)
	// slice operator
	fmt.Println(computer[3:5]) // 5 is excluded
	fmt.Println(computer[:4])  // 4 is excluded
	fmt.Println(computer[2:])  // from 2 to the last

	// Slices methods
	num1 := []int{1, 2, 3}
	num2 := []int{1, 2, 3, 4}
	fmt.Println(slices.Equal(num1, num2))   // return true if both are same
	fmt.Println(slices.Compare(num1, num2)) // return 0 if s1==s2, -1 if s1 < s2 , 1 if s1 > s2
}
