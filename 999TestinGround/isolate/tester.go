package main

import "fmt"

type student struct {
	fName string
	lName string
	yob   uint
}

func (st student) fullName() string {
	return st.fName + " " + st.lName
}

func main() {

	// var s2 student

	s1 := student{fName: "Karthik", yob: 1991, lName: "Nagadevara"}
	s2 := student{"Randy", "Paush", 1971}
	fmt.Println(s1.fullName())
	fmt.Println(s2.fullName())
}
