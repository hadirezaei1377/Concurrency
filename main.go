package main

// share memory in concurrency
// we want add 1000 times number to a variable by int type

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup

	myTestInt := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// pass pointer to function
		go incerement(&myTestInt, &wg)
	}

	wg.Wait()
	fmt.Println("End")
}

// recieve a pointer by type wg
// it need to be passed in main function
// pi = poniter int
func incerement(pi *int, wg *sync.WaitGroup) {
	// pass pi to an int
	i := *pi
	fmt.Println("value", i)
	*pi = i + 1
	wg.Done()
}
