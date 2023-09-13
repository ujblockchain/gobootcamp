package main

import "fmt"

func main() {
	// if statement will also perform a boolean check
	name := "john"

	//check if name is john
	if name == "john" {
		fmt.Printf("Welcome %v \n", name)
	}

	//more example
	isFinalYear := true

	//check if name is john
	if isFinalYear {
		fmt.Printf("Is John a final year student: %v \n", isFinalYear)
	}

	// if can also use else as a fall back,
	// just in case the if check condition falls,
	lastName := "Deo"

	if lastName == "Deo" {
		fmt.Printf("Welcome %v", lastName)
	} else {
		fmt.Println("Identity not verified")
	}

	//the if check can fail, the fall back will be else
	if lastName == "Joe" {
		fmt.Printf("Welcome %v \n", lastName)
	} else {
		fmt.Println("Identity not verified")
	}

	//you can also use multiple if checks, using else if
	nameChange := "Echo"

	if nameChange == "Doe" {
		fmt.Println("Welcome Doe")
	} else if nameChange == "James" {
		fmt.Println("Welcome James")
	} else if nameChange == "Echo" {
		fmt.Println("Welcome Echo")
	} else {
		fmt.Println("Identity not verified")
	}
}
