package main

import "fmt"

func main() {

	fmt.Println("Generic Typed returns")
	fmt.Println(StudentDetails("Karthik", 2013, 21100024))
	fmt.Println("Named Returns")
	fmt.Println(Minami("karthik", 29))
	// var numList = []float64{32, 141, 154, 11, 54, 111, 1788}
	// avg := MyAverage(numList)
	// Need to check why the above is not working.
	avg := MyAverage(32, 141, 154, 11, 54, 111, 1788)
	fmt.Printf("The average value of the is %v\n", avg)

}
