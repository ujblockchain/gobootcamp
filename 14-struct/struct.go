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
func (bio Bio) attendance() {
	fmt.Printf("%v was in %v today \n", bio.name, bio.school)
}

func main() {
	var person = Bio{
		name:   "Edmund Rotimi",
		school: "UJ",
	}

	checkAttendance(person)
}

// using struct with function
func checkAttendance(person Bio) {
	person.attendance()
}
