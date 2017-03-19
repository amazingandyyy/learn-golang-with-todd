package main

import "sort"
import "fmt"

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	s := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(s)
	fmt.Println("result: ", s)
	sort.Sort(sort.Reverse(s))
	fmt.Println("result 2: ", s)
}
