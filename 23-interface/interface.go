package main

import (
	"fmt"
)

// interface
type attendance interface {
	getAttendance()
}

// struct
type Bio struct {
	name   string
	school string
}

func (bio Bio) getAttendance() {
	fmt.Printf("%v was in school today \n", bio.name)
}

func main() {
	var person = Bio{
		name:   "Edmund Rotimi",
		school: "UJ",
	}
	//this works since person is a struct that has a method getAttendance
	checkAttendance(person)
}

// person parameter must have a method getAttendance
func checkAttendance(person attendance) {
	fmt.Printf("Hello")
}
