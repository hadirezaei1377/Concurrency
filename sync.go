package main

/*

import (
	"fmt"
	"sync"
	"time"
)

// using waitgroups to avoid goroutine death in order to sleep
var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	message("Hi User!")

	wg.Wait()
	fmt.Println("End")
}

func message(text string) {
	// how many goroutine we need ?
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go printMessage(text, i)
	}
}

func printMessage(text string, index int) {
	if index == 2 {
		time.Sleep(2 * time.Second)
	}
	fmt.Println(text, index)
	wg.Done()
}

*/
