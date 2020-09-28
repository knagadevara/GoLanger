package main

import (
	"fmt"
	"strconv"
)

var (
	ia         int = 50
	jb         int
	IamCorrect bool = true
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

	// type conversion example
	fmt.Printf("ia: %v , Type of ia: %T \n", ia, ia)
	jb = ia
	fmt.Printf("Assigning the value of ia to jb: %v , Type of jb: %T \n", jb, jb)
	fmt.Println("Converting an integer to float")
	var jb = float64(ia)
	fmt.Printf("jb: %v , Type of jb: %T \n", jb, jb)
	// Converting an integer to string
	var kc = strconv.Itoa(ia)
	fmt.Printf("kc: %v , Type of kc: %T \n", kc, kc)

	fmt.Printf("IamCorrect is fo type %T and contains %v\n", IamCorrect, IamCorrect)
}
