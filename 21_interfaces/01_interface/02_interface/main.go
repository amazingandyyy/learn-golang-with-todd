package main

import "fmt"

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(s shape) {
	// why do we need interface? not just instead using square type
	// using interface, we can have mutiple types under this interface
	fmt.Println("1", s)
	fmt.Println("2", s.area())
}

func main() {
	s := square{10}
	fmt.Printf("3 %T\n", s)
	info(s)
}
