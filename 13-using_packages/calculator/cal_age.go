package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalAge() {
	//init a reader to read from the terminal
	reader := bufio.NewReader(os.Stdin)

	//read user input for name
	fmt.Print("Enter your full name: ")
	name, _ := reader.ReadString('\n')

	//read user input for year
	fmt.Print("Enter current year: ")
	year, _ := reader.ReadString('\n')

	//read user input for dob
	fmt.Print("Enter your date of birth: ")
	DOB, _ := reader.ReadString('\n')

	//remove enter read with input
	cleanName := strings.Replace(name, "\n", "", -1)
	cleanYear := strings.Replace(year, "\n", "", -1)
	cleanDOB := strings.Replace(DOB, "\n", "", -1)

	//remove spaces in inputs
	cleanName = strings.TrimSpace(cleanName)
	cleanYear = strings.TrimSpace(cleanYear)
	cleanDOB = strings.TrimSpace(cleanDOB)

	// convert to inputs to numbers
	convYear, _ := strconv.ParseInt(cleanYear, 10, 64)
	convDOB, _ := strconv.ParseInt(cleanDOB, 10, 64)

	//calculate the age
	var age int64 = convYear - convDOB

	//print name and calculated age
	fmt.Printf("Your name is %v and you are %v years old", cleanName, age)
}
