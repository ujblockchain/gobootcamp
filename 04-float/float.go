package main

import "fmt"

func main() {
	// a variable is a placeholder, a container to hold something
	// "var" is used to declare variable

	// different ways to declare a float variable
	var burgerPrice = 24.57
	// var burgerPrice float64 = 24.56
	// var burgerPrice := 24.56

	//printing integer
	fmt.Println(burgerPrice)

	//different types of float
	// var maizePrice  float32  can take up to 6 number after the decimal point
	// var companyProfit  int16  can take up to 15 number after the decimal point

	// arithmetic operation with float
	// go will add extra zeros sometimes for precision sake.
	var shoppingCash float64 = 1900.80
	var currentPurchase float64 = 1789.60
	var changeLeft = shoppingCash - currentPurchase

	//
	fmt.Println(changeLeft)

	// addition and subtraction
	var schoolFee = 18780.90
	var registrationFee = 2345.90
	var registrationCost = schoolFee + registrationFee
	//
	fmt.Println(registrationCost)

	// Multiplication and division
	var walletAmount = 178.56
	var doubleAmount = walletAmount * 2
	//
	fmt.Println(doubleAmount)

	var changeAmount = 18.56
	var reduceAmount = changeAmount / 2
	fmt.Println(reduceAmount)
}
