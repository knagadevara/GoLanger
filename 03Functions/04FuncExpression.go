// Go does not allow multiple or nested functions, So,
// That leaves literally only one way to do nested functions
// it would be via assigning an anonymous function to a variable
// Which is known as func expressions.
// Decleration Syntax
// variableX := func() returnType { function Defination }
// Calling Syntax
// variableX() for non returns
// variableY = variableX() for returns

package main

import "fmt"

// SayYo demonstrates about the function expression usage
func SayYo() {

	// func expression, cannot be called out of the block
	SayHello := func() {
		fmt.Println("Hello")
	}

	// func expression. cannot be called out of the block
	SayMyName := func() string {
		return "Heisenberg!?"
	}
	SayHello()
	fmt.Printf("Type: %T \n", SayHello)
	i := SayMyName()
	fmt.Printf("Value of Name: %v \n", i)
	fmt.Printf("Type and returns: %T \n", SayMyName)
}

// As it is not possible to access the variables SayHello SayMyName out side of the function scope
// we would generally return the function which inturn returns the values to the variables.
// Re-designing SayMyName

// SayMyName is the function which returns a function which inturn returns a type of string
// This function is known as closure and can be accessable like any-other.
// decleration in main.go
// WhatsMyName := SayMyName()
// fmt.Printf("%v", WhatsMyName())
func SayMyName() func() string {

	return func() string {
		return "Hello I am HisenBurg!! \n"
	}
}

// WhatsYourAge returns a function of type int
func WhatsYourAge() func() int {

	return func() int {
		return 29
	}
}
