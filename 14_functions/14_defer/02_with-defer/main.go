package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}
func world2() {
	fmt.Println("world 2")
}

func main() {
	defer world2()
	hello()
	defer world()
}
