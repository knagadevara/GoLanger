package main

import "fmt"

//changeMyAge corrects the original age
func changeMyAge(ageAddress *int) int {

	fmt.Printf("Value: %v \t Type: %T \n", ageAddress, ageAddress)
	fmt.Println("Changing My age to original")
	*ageAddress = 30
	return *ageAddress
}

//PrintMyAge changes the original and prints
func PrintMyAge() {

	var agepaper int = 20
	fmt.Println("Before")
	fmt.Printf("Value: %v \t Type: %T \t Memory Value: %p \n", agepaper, agepaper, &agepaper)
	ActualAge := changeMyAge(&agepaper)
	fmt.Println("Changed in a different variable")
	fmt.Printf("Value: %v \t Type: %T \t Memory Value: %p \n", ActualAge, ActualAge, &ActualAge)
	fmt.Println("After Original Variable")
	fmt.Printf("Value: %v \t Type: %T \t Memory Value: %p \n", agepaper, agepaper, &agepaper)
}

// The Slice/Structs types would always change the original value rather than creating a new
// make is used to create a data structure with type, size, capacity

// var DetailsStack = make([]string, 10, 20)
// func changeMyName(DStack []string) {
// 	var MyName string = "Karthik"
// 	DStack[0] = MyName
// }
