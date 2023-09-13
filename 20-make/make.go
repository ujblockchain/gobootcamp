package main

import "fmt"

func main() {
	//create a list with capacity for 10,
	// make 2 space available during creation that can be filled
	myList := make([]string, 2, 10)
	myList[0] = " Hello"
	myList[1] = "Hi"
	//you can add more items since we have capacity for 10

	myList = append(myList, "13") // 7

	fmt.Printf("%v \n", myList)
	fmt.Printf("%v \n", myList[2])

}
