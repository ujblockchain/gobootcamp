package main

import "fmt"

func main() {
	//ways to declare a map
	var pcInfo = map[string]string{
		"Dell":     "Allienware",
		"HP":       "HP Omen",
		"Dynabook": "SatellitePro",
	}

	// pcInfo  := map[string]string{
	// 	"Dell":     "Allienware",
	// 	"HP":       "HP Omen",
	// 	"Dynabook": "SatellitePro",
	// }

	// var pcInfo map[string]string = map[string]string{
	// 	"Dell":     "Allienware",
	// 	"HP":       "HP Omen",
	// 	"Dynabook": "SatellitePro",
	// }

	//print values of a map using their key
	fmt.Printf("%v \n", pcInfo["Dell"])
	fmt.Printf("%v \n", pcInfo["HP"])
	fmt.Printf("%v \n", pcInfo["Dynabook"])

	//more on map
	var classPosition map[int16]string = map[int16]string{
		1: "John Deo",
		2: "Melvin Drum",
		3: "kelvin Smith",
	}

	fmt.Printf("%v \n", classPosition[1])
	fmt.Printf("%v \n", classPosition[2])
	fmt.Printf("%v \n", classPosition[3])

	//updating a map value using their keys
	classPosition[4] = "John Deo"
	fmt.Printf("%v \n", classPosition)

	//deleting the items from a map using their keys
	delete(classPosition, 4)
	fmt.Printf("%v \n", classPosition)

}
