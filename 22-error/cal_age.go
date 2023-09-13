package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//init a reader to read from the terminal
	reader := bufio.NewReader(os.Stdin)

	//read user input for name, we ignore the error for now
	fmt.Print("Enter your full name: ")
	name, err := reader.ReadString('\n')

	//check if there is an error
	if err != nil {
		fmt.Println("Invalid name")
		//stop the application
		panic(err)
	}

	//read user input for year, we ignore the error for now
	fmt.Print("Enter current year: ")
	year, err := reader.ReadString('\n')

	//check if there is an error
	if err != nil {
		fmt.Println("Invalid year entered")
		//stop the application
		panic(err)
	}

	//read user input for dob, we ignore the error for now
	fmt.Print("Enter your date of birth: ")
	DOB, err := reader.ReadString('\n')

	//check if there is an error
	if err != nil {
		fmt.Println("Invalid date of birth")
		//stop the application
		panic(err)
	}

	//remove enter read with input
	cleanName := strings.Replace(name, "\n", "", -1)
	cleanYear := strings.Replace(year, "\n", "", -1)
	cleanDOB := strings.Replace(DOB, "\n", "", -1)

	//remove spaces in inputs
	cleanName = strings.TrimSpace(cleanName)
	cleanYear = strings.TrimSpace(cleanYear)
	cleanDOB = strings.TrimSpace(cleanDOB)

	// convert to inputs to numbers
	convYear, err := strconv.ParseInt(cleanYear, 10, 64)

	//check if there is an error
	if err != nil {
		fmt.Println("Could not convert to number")
		//stop the application
		panic(err)
	}

	convDOB, err := strconv.ParseInt(cleanDOB, 10, 64)

	//check if there is an error
	if err != nil {
		fmt.Println("Could not convert to number")
		//stop the application
		panic(err)
	}

	//calculate the age
	var age int64 = convYear - convDOB

	//print name and calculated age
	fmt.Printf("Your name is %v and you are %v years old", cleanName, age)
}
