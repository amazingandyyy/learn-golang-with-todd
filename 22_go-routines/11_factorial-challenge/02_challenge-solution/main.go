package main

import (
	"fmt"
)

func main() {
	c := factorial(60)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		var total uint64 = 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
