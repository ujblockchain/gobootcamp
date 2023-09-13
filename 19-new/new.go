package main

import "fmt"

func main() {
	// new returns a pointer
	var age = new(int)
	randomAge := 45
	age = &randomAge

	//print the pointer address
	fmt.Println(age)
	//dereference the number
	fmt.Println(*age)

	//more example using a number
	var newNumber = new([]string)
	secondNewNumber := []string{"Hello"}
	newNumber = &secondNewNumber
	fmt.Println(newNumber)
}
