/* 🔥 Purpose of WaitGroup

- The primary purpose of a WaitGroup is to coordinate the synchronization of goroutines.
When you have multiple goroutines performing tasks concurrently, you may need to ensure that
all those tasks are completed before moving on to the next step in your program. A WaitGroup
helps you do that.


1) Create a WaitGroup: You start by creating a WaitGroup variable.
2) Add the number of goroutines to wait for: Before starting the goroutines, you call wg.Add(n),
										  where n is the number of goroutines you want to wait for.
3) Goroutines call Done() when finished: Each goroutine that is part of the WaitGroup should call
									  wg.Done() when it has finished its task. This
									  decrements the WaitGroup counter by one.
4) Wait for all goroutines to finish: The main goroutine or any other goroutine can call wg.Wait(),
								 which blocks until the counter inside the WaitGroup drops
								 to zero, meaning all goroutines have finished their tasks.

*/

package main

import (
	"fmt"
	"sync"
)

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing the task", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
}
