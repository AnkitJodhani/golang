package main

import "fmt"

func whoAmI(i interface{}) { // function will accept any kind of value like string, int, float, etc...
	switch cast := i.(type) {
	case int:
		fmt.Println("It's an Integer")
	case string:
		fmt.Println("It's a String")
	case bool:
		fmt.Println("It's a Boolean")
	case float32:
		fmt.Println("It's a float")
	default:
		fmt.Println("Not able to recognize", cast)
	}
}

func main() {
	whoAmI("passing string")
	whoAmI(232)
	whoAmI(232.33)
	whoAmI(true)
}
