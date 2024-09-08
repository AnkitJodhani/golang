package main

import (
	"fmt"
	"time"
)

func waitOneSec(done chan int) {
	time.Sleep(time.Second * 1)
	done <- 1
}

func waitTwoSec(done chan int) {
	time.Sleep(time.Second * 2)
	done <- 2
}

func main() {
	done := make(chan int)
	go waitOneSec(done)
	go waitTwoSec(done)
	fmt.Println(<-done)
}
