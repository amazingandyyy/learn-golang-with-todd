package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	if val, ok := myGreeting[7]; ok {
		delete(myGreeting, 7)
		fmt.Println("val: ", val)
		fmt.Println("ok: ", ok)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("ok: ", ok)
	}

	fmt.Println(myGreeting)
}
