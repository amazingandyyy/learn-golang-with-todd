package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("IntSlice: ", sort.IntSlice(n))
	sort.Ints(n)
	fmt.Println("result: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("result 2: ", n)
}
