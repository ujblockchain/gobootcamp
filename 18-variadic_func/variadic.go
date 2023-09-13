package main

import "fmt"

func main() {

	// variadic function
	numbers := []int64{23, 45, 67, 89, 90, 120}
	fmt.Println(numbers)

	//call in function as variadic
	addTwoNumber(numbers...)
	addThreeNumber(numbers...)

	//more examples
	newNumbers := []int64{23, 45}
	sum := addNumbers(newNumbers[0], newNumbers...)
	fmt.Printf("The sum is : %v", sum)

}

func addThreeNumber(numberThree ...int64) {
	fmt.Printf("%v", numberThree)
}

func addTwoNumber(numberTwo ...int64) {
	fmt.Println(numberTwo[0])
	fmt.Println(numberTwo[1])
}

func addNumbers(numberOne int64, numberTwo ...int64) int64 {
	var sum int64 = numberOne

	for _, value := range numberTwo {
		sum = sum + value
	}

	return sum
}
