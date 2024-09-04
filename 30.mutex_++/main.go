package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) incView(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views++ // critical section - race condition

}

func main() {
	var wg sync.WaitGroup
	ankitPost := post{views: 0}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go ankitPost.incView(&wg)
	}
	wg.Wait()
	fmt.Println(ankitPost.views)
}
