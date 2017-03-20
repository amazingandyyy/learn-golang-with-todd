package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			// after 1
			// no one is reciveing anymore
		}
	}()

	fmt.Println(<-c)
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?
