package main

import "fmt"

func main() {
	// a variable is a placeholder, a container to hold something
	// "var" is used to declare variable

	// different ways to declare a integer variable
	var age = 24
	// var age int64 = 24
	// var age := 24

	//printing integer
	fmt.Println(age)

	//different types of int
	// var numberInClass  int8  can take from -128 to 127
	// var numberInStadium  int16  can take from -32768 to 32767
	// var numberInSchool  int32  can take from -2147483648 to 2147483647
	// var numberInCountry int64  can take from -9223372036854775808 to 9223372036854775807

	// arithmetic operation with integers
	var dateOfBirth int64 = 1963
	var currentYear int64 = 2023
	var currentAge = currentYear - dateOfBirth

	//
	fmt.Println(currentAge)

	// addition and subtraction
	var numberOne = 23 + 89
	var numberTwo = 90 - 78
	//
	fmt.Println(numberOne)
	fmt.Println(numberTwo)

	// you can also do
	var distanceToJoburg = 900
	var distanceToCapeTown = 5000
	//
	var distanceDifference = distanceToCapeTown - distanceToJoburg
	//
	fmt.Println(distanceDifference)

	// Multiplication and division
	var numberThree = 23 * 89
	var numberFour = 90 - 2
	//
	fmt.Println(numberThree)
	fmt.Println(numberFour)

	// you can also do
	var presentAge = 34
	var ageDouble = presentAge * 2
	var ageReduce = presentAge / 2
	//
	fmt.Println(ageDouble)
	fmt.Println(ageReduce)

	//using modulus "%": this returns a reminder of division
	var reminderCal = 23 % 2
	fmt.Printf("The reminder is: %v \n", reminderCal)

	//more one modulus "%": this returns a reminder of division
	var reminderCalTwo = 20 % 2
	fmt.Printf("The reminder is: %v \n", reminderCalTwo)

}
