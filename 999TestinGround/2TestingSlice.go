package main

import "fmt"

var StdLst = make([][]string, 3)
var student1 = make([]string, 3)
var student2 = make([]string, 3)
var student3 = make([]string, 3)

func main() {

	//	Sending Values in student1
	student1[0] = "Karthik"
	student1[1] = "Nagadevara"
	student1[2] = "17-JAN-1991"

	student2[0] = "Srinivas"
	student2[1] = "Palaka"
	student2[2] = "23-JAN-1985"

	student3[0] = "Randy"
	student3[1] = "Paush"
	student3[2] = "21-DEV-1971"

	StdLst = append(StdLst, student1, student2, student3)

	fmt.Println(StdLst)

}
