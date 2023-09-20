package main

import "fmt"

func main() {
	// a variable is a placeholder, a container to hold something
	// "var" is used to declare variable

	// different ways to declare a variable
	var name = "John Deo"
	// var name string = "John Deo"
	// name := "John Deo"

	//you can also print a variable using fmt Print
	//fmt.Print(name)
	fmt.Println(name)

	//interpolation using fmt Printf
	fmt.Printf("My name is : %v ", name)
	//you can "\n" for new line
	fmt.Printf("My name is : %v \n", name)

	//declaring variable name using camelCase
	var departmentName = "Philosophy"
	fmt.Printf("My department is : %v \n", departmentName)

}
