package main

import "fmt"

func main() {
	fc := factorial(60)
	for n := range fc {
		fmt.Println("Total:", n)
	}
}

func factorial(n uint64) chan uint64 {
	r := make(chan uint64)
	go func() {
		var result uint64 = 1
		for i := n; i > 0; i-- {
			result *= i
		}
		r <- result
		close(r)
	}()
	return r
}
