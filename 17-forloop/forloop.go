package main

import "fmt"

func main() {

	/// first type of for loop: for loping through a list

	var myTestScore []int64 = []int64{89, 90, 91, 100} ///

	//hoe items in a list are numbered
	// // 0 = 89
	// // 1 = 90
	// // 2 = 91
	// // 3 = 100
	// // LEN = 4 //LEN -1

	var sum int64 = 0

	// requirements for a for loop
	// // DECELERATION
	// // CONDITION
	// // COUNTER

	for i := 0; i < len(myTestScore); i++ {
		sum = sum + myTestScore[i]

		////// at i= 0, sum 0 + myTestScore[0] => 89
		///// at i=1, sum = sum + myTestScore[1] => 89 + 90 => 179
		////// at i = 2, sum = sum + myTestScore[2] => 179+ 91 =>270
		////// at i = 3, sum = sum + myTestScore[3] => 279 + 100 => 370
		////// at i = 4, false
		///end
	}

	fmt.Println(sum)

	// second type of for loop: used for conditions
	// it keeps running as long as the condition is true

	var isAdult bool = true
	var age int32 = 10

	for isAdult {
		fmt.Println(age)

		if age < 18 {
			isAdult = false
		}
	}

	// third type of for loop; used with maps and list also

	// with makes
	var phoneMaker map[string]string = map[string]string{
		"Google":  "Pixel 7a",
		"Samsung": "Z flip",
		"Itel":    "P40",
		"Iphone":  "Iphone14",
	}

	for key, value := range phoneMaker {
		fmt.Printf("Key: %v : value: %v \n", key, value)
	}

	//with list
	var testScore []int64 = []int64{89, 90, 91, 100}
	for index, value := range testScore {
		fmt.Printf("Value: %v : value: %v \n", index, value)
	}

}
