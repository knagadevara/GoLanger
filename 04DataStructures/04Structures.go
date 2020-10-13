package main

import "fmt"

type student struct {
	fName string
	lName string
	yob   uint
}

type scienceYR struct {
	student
	IsEngineer bool
}

func (st student) fullName() string {
	return st.fName + " " + st.lName
}

func main() {

	// var s2 student

	s1 := student{fName: "Karthik", lName: "Nagadevara", yob: 1991}
	s2 := student{"Randy", "Paush", 1971}

	eng1 := scienceYR{student: s1, IsEngineer: true}
	eng2 := scienceYR{
		student: student{
			fName: "Jaikanth",
			lName: "Nagre",
			yob:   1965,
		},
		IsEngineer: false,
	}
	fmt.Println(eng1.fullName())
	fmt.Println(s2.fullName())
	fmt.Println(eng2.fullName())
}
