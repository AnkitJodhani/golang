package main

import "fmt"

// any | string | int | interface{} | comparable

func printFunc[T int | string](slice []T) { // only accept string or integer

	for _, value := range slice {
		fmt.Println(value)
	}
}
func printFuncTest[T comparable](slice []T) { // comparable is interface which has all the type

	for _, value := range slice {
		fmt.Println(value)
	}
}

type Stack[T string | int] struct {
	element []T
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"Golang", "JavaScript"}

	printFunc(nums)
	printFunc(names)
	printFuncTest(nums)
	printFuncTest(names)

	s1 := Stack[int]{
		element: []int{1, 23, 3},
	}
	s2 := Stack[string]{
		element: []string{"ankit", "piyush"},
	}
	fmt.Println(s1)
	fmt.Println(s2)

}
