package main

import "fmt"

func makeGreeter() func() string {
	// return func() string {
	// 	return "Hello world!"
	// }
	say := func() string {
		return "Hello world!"
	}
	return say
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
