package main

import "fmt"

// we can return multiple values in golang
func listLanguage() (string, string, string, string) {
	return "golang", "python", "javascript", "java"
}

func main() {
	// ---------- functions ------------
	l1, l2, l3, _ := listLanguage()
	fmt.Println(l1, l2, l3)
}
