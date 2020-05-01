package main

import "fmt"

var forTrunc float64 = 10.2
var intNum int

func truncateInt(y float64) int {
	var x = int(y)
	return x
}

func main() {
	fmt.Println("Step 1")

	intNum = truncateInt(forTrunc)
	fmt.Println(intNum)

	fmt.Println("/////////")

	fmt.Println("Step 2")

	forTrunc = 6.2
	intNum = truncateInt(forTrunc)
	fmt.Println(intNum)
}
