package main

import "fmt"

func main() {
	// pointers return a memory address
	var newNumber *int64
	var secondNewNumber int64 = 56
	newNumber = &secondNewNumber

	//print address
	fmt.Println(newNumber)

	//dereference to get back 56
	fmt.Println(*newNumber)

}


