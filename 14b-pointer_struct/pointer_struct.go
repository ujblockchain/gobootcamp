package main

import (
	"fmt"
)

// struct type
type Bio struct {
	name   string
	school string
}

// struct method
func (bio *Bio) attendance() {
	fmt.Printf("%v was in %v today \n", bio.name, bio.school)
}

func main() {
	//using pointer with struct
	var person = Bio{
		name:   "Edmund Rotimi",
		school: "UJ",
	}

	// function accept a pointer
	// user & to convert to pointer
	myName, school := checkAttendance(&person)
	fmt.Printf("Your name is : %v \n", myName)
	fmt.Printf("Your school is : %v \n", school)

	//pointer can be dereference also by you
	newPerson := &person

	//`&` is used for dereferencing
	fmt.Printf("Your name is : %v \n", (*newPerson).name)
	fmt.Printf("Your school is : %v \n", (*newPerson).school)

	//go can dereference automatically for You
	fmt.Printf("Your name is : %v \n", newPerson.name)
	fmt.Printf("Your school is : %v \n", newPerson.school)

	//call a method
	//go dereference automatically for You
	person.attendance()
}

// using struct with function
func checkAttendance(person *Bio) (string, string) {
	//go will automatically dereference
	return person.name, person.school
}
