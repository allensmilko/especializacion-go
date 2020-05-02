package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var forTrunc float64
var intNum int

func truncateInt(y float64) int {
	var x = int(y)
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Step 1")
	fmt.Println("Enter first float value : ")
	scanner.Scan()
	input := scanner.Text()
	forTrunc, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	}
	intNum = truncateInt(forTrunc)
	fmt.Println(intNum)

	fmt.Println("/////////")

	fmt.Println("Step 2")
	fmt.Println("Enter second float value : ")
	scanner.Scan()
	input = scanner.Text()
	forTrunc, err = strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	}
	forTrunc = 6.2
	intNum = truncateInt(forTrunc)
	fmt.Println(intNum)
}
