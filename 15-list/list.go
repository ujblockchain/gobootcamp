package main

import "fmt"

func main() {
	//ways to declare a dynamic list
	var myList []string = []string{"rice", "apples", "tomatoes", "chilly", "corn"}
	//var myList  = []string{"rice", "apples", "tomatoes", "chilly", "corn"}
	//myList := []string{"rice", "apples", "tomatoes", "chilly", "corn"}

	// accessing the items in the list, it starts from 0
	fmt.Printf("The first item on my list is %v \n", myList[0])
	fmt.Printf("The second item on my list is %v \n", myList[1])
	fmt.Printf("The third item on my list is %v \n", myList[2])

	// getting the length of the list
	fmt.Printf("The length of my list is %v \n", len(myList))

	// get last item in the list
	fmt.Printf("The last item on my list is %v \n", myList[len(myList)-1])

	//list with fixed length
	var favPcList [4]string = [4]string{"Alienware", "MSI", "Asus", "Dynabook"}
	fmt.Printf("%v \n", favPcList)

	//more example of fixed length list, but initialize first
	var favSubject [6]string //0 1 2 3 4 5
	//set items to the list
	favSubject[0] = "Physics"
	favSubject[1] = "Chemistry"
	favSubject[2] = "Maths"
	favSubject[3] = "Computer Science"
	favSubject[4] = "Biology"
	favSubject[5] = "Economics/Geography/English"
	//print the items in the list
	fmt.Printf("The first item on my list is %v \n", favSubject[0])
	fmt.Printf("The second item on my list is %v \n", favSubject[1])
	fmt.Printf("The third item on my list is %v \n", favSubject[2])

	//more example of a list with number
	//length starts counting from 1 eg len(oddNumberList) = 6
	// index start counting 0 eg. oddNumberList[0] => 1 =. 0, 1, 2, 3, 4, 5, 6
	var oddNumberList []int64 = []int64{1, 3, 5, 7, 9, 11} /// 6:
	fmt.Printf(" %v \n", oddNumberList)

	// using append to add to a list
	// append can is append to a new variable not the original
	updatedList := append(oddNumberList, 13) // 7

	fmt.Printf("Old List: %v \n", oddNumberList)
	fmt.Printf("New List: %v \n", updatedList)

	//List operation: explaining slice, capacity and length
	var randomFloatNumbers []float64 = []float64{12.34, 23.45, 45.78, 90.8, 0.23}
	shortFloatList := randomFloatNumbers[0:2] // 1-1 //[12.34, 23.45]

	//copy the whole list using slice
	secondShortFloatList := randomFloatNumbers[0:]

	//using concept of capacity to access value to the right
	thirdFloatList := shortFloatList[2:5]
	fourthShortFloatList := shortFloatList[:5] //4-1=3

	fmt.Printf("%v \n", shortFloatList)
	fmt.Printf("len: %v capacity: %v\n", len(shortFloatList), cap(shortFloatList))

	fmt.Printf("%v \n", secondShortFloatList)
	fmt.Printf("len: %v capacity: %v\n", len(secondShortFloatList), cap(secondShortFloatList))

	fmt.Printf("%v \n", thirdFloatList)
	fmt.Printf("%v \n", fourthShortFloatList)
}
