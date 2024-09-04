package main

import (
	"fmt"
	"sync"
	"time"
)

func incNumber() func() int {
	num := 0
	return func() int {
		var mu sync.Mutex
		mu.Lock()
		num += 1
		mu.Unlock()
		fmt.Println(num)
		return num
	}
}

func main() {
	doingSomething := incNumber()
	for i := 0; i < 100; i++ {
		go doingSomething()
	}

	time.Sleep(time.Second * 2)
}
