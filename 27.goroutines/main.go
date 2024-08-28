package main

import (
	"fmt"
	"time"
)

func doingSomething(id int) {
	fmt.Println("Executing Task ", id)

}

func main() {
	for i := 0; i < 10; i++ {
		go doingSomething(i) // "go" is keyword
	}

	// wait for 1 second
	time.Sleep(time.Second * 1) // preventing main function from poping from STACK
}
