package main

import (
	"fmt"
	"os"
)

func main() {
	f, error := os.Open("31.files/readme.txt")
	if error != nil {
		// log the error
		// fmt.Println("Got the error", error)
		panic(error)
	}

	fileInfo, error := f.Stat()

	fmt.Println(fileInfo.Name())    // return name
	fmt.Println(fileInfo.IsDir())   // return true or false
	fmt.Println(fileInfo.ModTime()) // return modification time
	fmt.Println(fileInfo.Size())    // size of the file in Bytes
	fmt.Println(fileInfo.Mode())    // return the permission
}
