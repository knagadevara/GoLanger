package main

import "fmt"

var testX int

func PrintMyName(name string) func() string {
	return func() string {
		return fmt.Sprintf("Hello %s", name)
	}
}

func incrementX() int {
	testX++
	return testX
}

func incrementY() func() int {
	var testY int
	return func() int {
		testY++
		return testY
	}
}

func main() {

	var getName func() string = PrintMyName("SivaRam")

	fmt.Printf("Type: %T \n", getName)
	fmt.Printf("Type: %T \n", getName())
	fmt.Printf("Value %v \n", getName())
	fmt.Printf("X %d \n", incrementX())
	incrementX()
	incrementX()
	incrementX()
	fmt.Printf("X %d \n", incrementX())
	fmt.Printf("Printing X directly %d \n", testX) //Can still access X
	fmt.Printf("Y %d \n", incrementY()())          // A direct execution will result in a
	//new function execution everytime we run
	fmt.Printf("Y %d \n", incrementY()())
	fmt.Printf("Y %d \n", incrementY()())
	fmt.Printf("Y %d \n", incrementY()()) // this functio will return 1 always

	incrementStaticAssignmentY := incrementY() // Assigning the memory address to a variable,
	//making it a function expression
	incrementStaticAssignmentY()
	incrementStaticAssignmentY()
	incrementStaticAssignmentY()
	fmt.Printf("Y %d \n", incrementStaticAssignmentY()) // the state of the function is presisted
	//and now it gets incremented, but Y is not accessable from outside

}
