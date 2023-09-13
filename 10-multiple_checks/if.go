package main

import (
	"fmt"
)

func main() {
	//user input
	numberAge := 78

	//check if age is greater than 18 and less tha 60 years only
	if numberAge >= 18 && numberAge > 60 {
		fmt.Println("Welcome to BS Club, go to row 1")
	} else {
		fmt.Println("You are not allowed into the club")
	}

	//more examples using `and`
	if numberAge == 18 {
		fmt.Println("Welcome to BS Club, go to row 1")
	} else if numberAge > 18 && numberAge < 30 { //check if age is greater 18 and less than 30
		fmt.Println("Welcome to BS Club, go to row 2")
	} else if numberAge >= 30 && numberAge < 60 { //check if age is greater or equal to 30 and less than 60
		fmt.Println("Welcome to BS Club, go to row 3")
	} else if numberAge >= 60 { //check if age is greater or equal to 60
		fmt.Println("Welcome to BS Club, go to vip row")
	} else {
		fmt.Println("You are not allowed into the club")
	}

	//more examples using `or`
	newNumberAge := 56
	if newNumberAge == 18 {
		fmt.Println("Welcome to BS Club, go to row 1")
	} else if newNumberAge > 18 || newNumberAge < 30 { //check if age is greater 18 and less than 30
		fmt.Println("Welcome to BS Club, go to row 2")
	} else {
		fmt.Println("You are not allowed into the club")
	}
}
