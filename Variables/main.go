package main

import "fmt"

func main() {
	fmt.Println("Play with variables")

	// Basic way to declare a variable
	var name string = "guneet"
	var num int = 1

	fmt.Println(name)
	fmt.Println(num)

	// another way (It automaticaaly check the value and make the variable of that type)
	var anotherName = "guneet"
	var anotherNum = 2

	fmt.Println(anotherName)
	fmt.Println(anotherNum)

	// multiple varibale declaration
	var first, last string = "guneet", "singh"
	fmt.Println(first, last)

	// Smart way
	smartName := "Hello Guneet"

	// Let's Check type of variable
	fmt.Println(smartName)
	fmt.Printf("Type is: %T", smartName)
}
