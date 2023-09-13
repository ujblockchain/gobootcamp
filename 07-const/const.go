package main

import "fmt"

func main() {
	// a variable is a placeholder, a container to hold something
	// "var" is used to declare variable

	// declaring a const(constant) variable
	const studentNumber = 22127654
	//const studentNumber int = 22127654

	// you can not change a const variable after it has been defined
	//studentNumber = 45

	//printing
	fmt.Println(studentNumber)

	//more example
	const fullName = "John Deo"

	//printing
	fmt.Println(fullName)

	//using iota with const
	const (
		firstNumber = iota
		secondNumber
		thirdNumber
	)

	fmt.Println(firstNumber)
	fmt.Println(secondNumber)
	fmt.Println(thirdNumber)

	//using iota but start counting from one
	const (
		firstPosition = iota + 1
		secondPosition
		thirdPosition
	)

	fmt.Println(firstPosition)
	fmt.Println(secondPosition)
	fmt.Println(thirdPosition)

}
