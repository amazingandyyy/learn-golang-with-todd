package main

import "fmt"

func main() {
	fmt.Printf("decimal \t binary \t hexadecimal \t UFT-8 \t Type \t Value\n")
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d \t \t %b \t \t %x \t \t %q \t %T \t %v \n", i, i, i, i, i, i)
	}

	// fmt.Printf("%q \n", 65)
}
