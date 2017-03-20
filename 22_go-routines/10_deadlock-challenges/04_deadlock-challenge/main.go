package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c)
		// hanging there waiting for the one after 9
		// but doesn't receive any after 9
	}
}

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
// Why are all goroutines asleep?
// How can we fix this?
