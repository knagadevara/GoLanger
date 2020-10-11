package main

import "fmt"

// Map is reference type very similar to dictionary and it points to underlying hash_Table
//Syntax: var map_IdemtefierName = make(map[Key_Type]Value_Type , length)
// There is no append for map
var mapTest = make(map[string]int, 10)

func main() {

	mapTest["Karthik"] = 1991
	mapTest["Srinivas"] = 1985
	mapTest["Randy"] = 1971
	mapTest["Shashi"] = 1995
	mapTest["Sharan"] = 2000

	fmt.Println("map: ", mapTest)
	fmt.Println(mapTest["Randy"])

	e11, e12 := mapTest["Randiee"]
	fmt.Println(e11, e12, len(mapTest))

	// deleting an entry in the map
	delete(mapTest, "Randy")
	fmt.Println(mapTest)

	// to loop over the map

	for key, value := range mapTest {

		fmt.Println(key, " - ", value)
	}
}
