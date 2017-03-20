package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	// no one is sycly waiting to recieve the last one
	fmt.Println(<-c) // we need to run this go when we asign 1 to the c channel
}

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?
