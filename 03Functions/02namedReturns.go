package main

import "fmt"

//Minami Exports to 01multiReturns.go
func Minami(name string, age int8) (s string, a int8) {
	a = age
	s = fmt.Sprint(name)
	return
}

// Function returns
// Named returns

// Syntax:
// func RECIEVER <FunctionName>(parameter1-name parameter-type , parameter2-name parameter-type) (return-name1 return-type , return-name2 return-type) {
// return-name1 = parameter1-name
// return-name2 = parameter2-name
// }
