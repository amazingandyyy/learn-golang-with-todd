package main

import "fmt"

func zero(z *int) { // argument is in type of address of where the int is stored, called int pointer
	*z = 0 // sign(change) the value of that address to 0, dereference, unviel the value of the address
	// open the letter and change the content and put it back to the mail box
}

func main() {
	x := 5
	zero(&x)       // passing address of where the 5 is stored
	fmt.Println(x) // x is 0
}
