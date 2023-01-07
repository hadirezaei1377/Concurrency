package main

import (
	"fmt"
	"time"
)

//   // one core , one thread , one goroutine
// func main() {
// 	message()
// }

// func message() {
// 	fmt.Println("Hi User!")
// }

// division of this to two goroutines
// main go routine work alot and has main tasks and another goroutine has small task
func main() { // main goroutine is working
	go message() // now bey keyword go before function , create a new goroutine
	time.Sleep(time.Second)
}

func message() {
	fmt.Println("Hi User!")
}

/*
main goroutine is created and working , while its working , its arrives to new goroutine and start it.
main after does its tasks , dies and doesnt stop for new goroutine and close the proggram without any output.
we should time sleep for main for freezing it that the other goroutine works.
*/
