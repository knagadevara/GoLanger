package main

import (
	"fmt"
)

// StudentDetails var Studentstruct
func StudentDetails(name string, Yoj int16, Htno int32) (string, int32) {
	return fmt.Sprintf("Name: %s \t Year: %d \n", name, Yoj), Htno
}
