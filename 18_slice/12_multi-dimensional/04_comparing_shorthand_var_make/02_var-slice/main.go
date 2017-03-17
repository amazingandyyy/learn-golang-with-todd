package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	student[0] = "Todd" // error
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
