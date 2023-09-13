package main

import (
	"fmt"
)

func main() {
	name := "John Deo"

	//call function with argument
	checkAttendance(name)

	//profile
	university := "University of Johannesburg"

	//call function with two arguments
	myBio(name, university)

	//function with return 
	sum := addTwoNumbers(78, 89)
	fmt.Println(sum)
	
	//two return in function
	myName, age := nameAge("James Mel", 45)
	fmt.Printf("My name is : %v \n", myName)
	fmt.Printf("My age is : %v \n", age)

}

// function with a parameter
func checkAttendance(personName string) {
	fmt.Printf("%v is in the meeting today \n", personName)
}

// function with two or more parameter
func myBio(name string, university string) {
	fmt.Printf("My name is %v and I am studying in %v \n", name, university)
}

//using return in function
func addTwoNumbers(numberOne int64, numberTwo int64) int64{
	sum := numberOne + numberTwo
	return sum
}

//returning two or more
//function simply returns name and age
func nameAge(name string, age int32) (string, int32){
	return name, age
}

