package HelloWorld

import (
	"fmt"
	"strconv"
)

var (
	ia         int = 50
	jb         int
	IamCorrect bool = true
)

func HelloWorld() {

	fmt.Println("Hello, World!! I am GO")
	fmt.Println("Variable assignment")
	var a int
	a = 10
	var b int = 20
	c := 30
	fmt.Println("'var a int' \n 'a = 10' \n 'var a int = 10' \n 'a := 10' are Similar")
	fmt.Println(a, b, c)

	// knowing the type of a variable
	fmt.Printf("ia is of Type: %T and it contains %v\n", ia, ia)
	fmt.Printf("IamCorrect is of Type: %T and it contains %v\n", IamCorrect, IamCorrect)
	jb = ia
	fmt.Printf("Assigning the value of ia to jb: %v , Type of jb: %T \n", jb, jb)
	fmt.Println("Converting an integer to float float64(ia)")
	var jb = float64(ia)
	fmt.Printf("jb: %v , Type of jb: %T \n", jb, jb)
	fmt.Println("Converting an integer to string var kc = strconv.Itoa(ia)")
	var kc = strconv.Itoa(ia)
	fmt.Printf("kc: %v , Type of kc: %T \n", kc, kc)
}
