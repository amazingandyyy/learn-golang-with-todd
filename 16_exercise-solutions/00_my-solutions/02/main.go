package main

import (
	"fmt"
)

var take = func(a int) (float64, bool) {
	devides := float64(a) / 2
	evenOrNot := a%2 == 0
	return devides, evenOrNot
}

func main() {
	fmt.Println(take(7))
}
