package main

// share memory in concurrency
// we want add 1000 times number to a variable by int type

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Value int
}

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup

	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// pass pointer to function
		go incerement(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("value:", counter.Value)
	fmt.Println("End")
}

// recieve a pointer by type wg
// it need to be passed in main function
// pi = poniter int
func incerement(counter *Counter, wg *sync.WaitGroup) {
	counter.Lock()
	// give value
	i := counter.Value
	// pass pi to an int

	counter.Value = i + 1

	wg.Done()
	// finish processing
	counter.Unlock()
}
