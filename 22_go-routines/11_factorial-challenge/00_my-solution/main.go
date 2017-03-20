package main

import "fmt"

var fc = make(chan uint64)

func main() {
	go func() {
		factorial(50)
	}()
	for n := range fc {
		fmt.Println("Total:", n)
	}
}

func factorial(n uint64) {
	var result uint64 = 1
	for i := n; i > 0; i-- {
		result = result * i
	}
	fc <- result
	close(fc)
}
