package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var word string

func finded(y string) bool {
	x := strings.ContainsAny(y, "i|a|n")
	return x
}

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first text: ")
	word, _ = input.ReadString('\n')
	find := finded(word)
	fmt.Println(find)

	if find {
		fmt.Println("Word 1")
		fmt.Println("Found!")
		fmt.Sprintf("The word is: %s", word)

	} else {
		fmt.Println("Not Found!")
		fmt.Sprintf("The word is: %s", word)
	}

	fmt.Println("///////////////")

	fmt.Print("Enter second text: ")
	word, _ = input.ReadString('\n')
	find = finded(word)
	fmt.Println(find)
	if find {
		fmt.Println("Word 2")
		fmt.Println("Found!")
		fmt.Sprintf("The word is: %s", word)

	} else {
		fmt.Println("Not Found!")
		fmt.Sprintf("The word is: %s", word)
	}
}
