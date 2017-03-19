package main

import (
	"fmt"
	"sort"
)

// Resource: https://godoc.org/sort

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	// ~=> sort.Sort(sort.StringSlice(s))
	// ~=> sort.StringSlice(s).Sort()
	fmt.Println("result: ", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	// StringSlice will give us Len, Swap, Less three methods
	// StringSlice attaches the methods of Interface to []string
	// It doesn't return a Interface interface
	fmt.Println("result 2: ", s)
}
