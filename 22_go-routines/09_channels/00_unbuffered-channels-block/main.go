package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			// fmt.Println("1", c)
			// fmt.Println("2", <-c)
		}
	}()

	go func() {
		for {
			fmt.Println("3", <-c)
		}
	}()

	time.Sleep(time.Millisecond)
}
