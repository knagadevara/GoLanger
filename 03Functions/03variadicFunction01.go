package main

import "fmt"

// CalculateAverage External available function has the capablities
// '...float64' as paramerter enables the functionality of taking unlimited
// number of arguments of type float64 .looping over through range function which takes in items and returns index, value
func CalculateAverage(numList ...float64) (avg float64) {
	fmt.Println(numList)
	fmt.Printf("Type of variable: %T \n", numList)
	var total float64

	for _, numListValue := range numList {
		total += numListValue
	}
	return total / float64(len(numList))
}

func main() {

	someList : make([]int , 0)
	someList = [5,10,31,40,32]

}
