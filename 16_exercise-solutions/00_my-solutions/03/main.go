package main

import (
	"fmt"
)

func largest(n ...int) int {
	var result int
	for _, e := range n {
		if e > result {
			result = e
		}
	}
	return result
}

func main() {
	fmt.Println(largest(1, 2, 3, 4, 5, 6, 100))
}
