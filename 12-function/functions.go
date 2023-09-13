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
}

// function with a parameter
func checkAttendance(personName string) {
	fmt.Printf("%v is in the meeting today \n", personName)
}

// function with two or more parameter
func myBio(name string, university string) {
	fmt.Printf("My name is %v and I am studying in %v \n", name, university)
}
