package main

import "fmt"

func main() {
	// a variable is a placeholder, a container to hold something
	// "var" is used to declare variable

	// different ways to declare a rune variable
	var countryCurrency = '€'
	// var countryCurrency rune = '€'
	// countryCurrency := '€'

	//printing rune storage number
	fmt.Println(countryCurrency)

	//printing rune value
	fmt.Println(string(countryCurrency))

	//more example
	favoriteEmoji := '😇'

	//printing rune storage number
	fmt.Println(favoriteEmoji)

	//printing rune value
	fmt.Println(string(favoriteEmoji))

}
