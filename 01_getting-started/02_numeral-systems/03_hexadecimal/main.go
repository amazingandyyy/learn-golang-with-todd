package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \t <- Hexadeciaml of 42 without leading zeros \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \t <- Hexadeciaml of 42 with leading zeros \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#o \t <- Octal(8 base) of 42 with leading zeros \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \t <- Hexadeciaml of 42 with leading zeros \n", 42, 42, 42)
}
