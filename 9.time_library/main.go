package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time library")
	fmt.Println(time.Now())           // return current time
	fmt.Println(time.Now().Weekday()) // return name of the day
	fmt.Println(time.Monday)          // return current date
}
