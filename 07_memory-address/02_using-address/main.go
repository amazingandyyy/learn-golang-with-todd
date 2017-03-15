package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters, meters2 float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	fmt.Print("Enter meters swam2: ")
	fmt.Scan(&meters2)
	yards := meters * metersToYards
	yards2 := meters2 * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
	fmt.Println(meters2, " meters is ", yards2, " yards.")
}
