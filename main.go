package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, World!! I am GO")
	fmt.Println("Variable assignment")
	var a int
	a = 10
	var b int = 20
	c := 30

	fmt.Println("'var a int' \n 'a = 10' \n is similar to \n'var a int = 10' \nis similar to \n'a := 10'")
	fmt.Println(a, b, c)
}

//### running "go run main.go"
//### building "go build <path to main.go>"
// := can only be used while declaring a new variable
// All the variables can be declared at the package level inside var block
