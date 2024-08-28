package main

import "fmt"

func main() {

	var student map[string]string
	fmt.Println(student)

	var teacher = map[string]bool{"Megha": true}
	fmt.Println(teacher)

	myMap := map[int]bool{1: true, 2: false}
	fmt.Println(myMap)

	employee := make(map[string]string)
	fmt.Println(employee)

}
