package main

// MyAverage External Calling demonstrated the capablities of taking unlimited number of arguments of type float64
// As '...float64' enables the functionality paramerter defination
// looping over through range function which takes in items and returns index, value
func MyAverage(numList ...float64) float64 {
	// fmt.Println(numList)
	// fmt.Printf("Type of variable: %T \n", numList)
	var total float64

	for _, numListValue := range numList {
		total += numListValue
	}
	return total / float64(len(numList))
}
