package main

import (
	"fmt"
)

func main() {
	// var x [256]string

	// fmt.Println(len(x))
	// fmt.Println(x[0])
	// for i := 0; i < 256; i++ {
	// 	x[i] = string(i)
	// }
	// for _, v := range x {
	// 	fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	// }
	var bb = string(127)
	fmt.Println([]byte(bb))
	fmt.Printf("%b \n", []byte(bb))
}
