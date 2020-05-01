package main

import "fmt"

var greeting string

type User struct {
	Name string
	Age  int
}

func greetUser(greet string, name string, age int) string {
	user := User{Name: name, Age: age}
	greeting = fmt.Sprintf("%s : %s, you have %d old! ", greet, user.Name, user.Age)
	return greeting
}

func main() {
	var greet string = "Hello"
	var name string = "Jeison"
	var age int = 27

	fmt.Print(greetUser(greet, name, age))
}
