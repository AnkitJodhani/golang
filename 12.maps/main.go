package main

import (
	"fmt"
	"maps"
)

func main() {
	// 🔥 maps -> similar to dictionary in python

	student := make(map[string]string)

	student["name"] = "Ankit Jodhani"
	student["fname"] = "Hasmukhbhai Jodhani"
	fmt.Println(student)
	fmt.Println(student["name"])
	fmt.Println(student["fname"])
	fmt.Println(student["keyDoesNotExist"] == "") // return true

	marks := make(map[string]int)
	marks["january"] = 10
	marks["february"] = 20
	marks["march"] = 30
	marks["april"] = 40
	fmt.Println(marks)
	fmt.Println(marks["january"])
	fmt.Println(marks["keyDoesNotExist"]) // return zero

	teacher := make(map[string]bool) // declare & initilize with NO KEY
	// var teacher map[string]bool   // declare
	// var teacher = map[string]bool{"key1": true} // declare & initilize with key-value
	teacher["key1"] = true
	fmt.Println(teacher)
	fmt.Println(teacher["key1"])
	fmt.Println(teacher["keyDoesNotExist"])

	school := make(map[string]string)

	school["name1"] = "Sardar Patel"
	school["name2"] = "Ashadeep Vidhyalaya"
	school["name3"] = "Primary School"
	school["name4"] = "Science School"
	school["name5"] = "Indian Institue of Technology"
	fmt.Println(school)
	delete(school, "name3") // delete an element
	fmt.Println(school)
	clear(school)
	fmt.Println(school)

	// 🔥 declare with initilization
	university := map[string]int{"PPSU": 90, "UTU": 80}
	fmt.Println(university)

	// 🔥🔥🔥🔥 Okay syntax -> 💎
	// first value is VALUE that key hold & second value is boolean true or false
	value, ok := university["DTU"]
	fmt.Println(value)
	fmt.Println(ok)

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value does not exist")
	}

	// maps method
	m1 := map[string]string{"name": "ankit", "fname": "hasubhai"}
	m2 := map[string]string{"name": "ankit", "fname": "hasubhai"}

	fmt.Println(maps.Equal(m1, m2))

}
