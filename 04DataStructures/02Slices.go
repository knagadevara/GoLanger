package main

import "fmt"

var mySlice = make([]uint, 0, 7)
var i uint

// Slice, Maps, Channels are reference types to the underlaying premitive Data Structure of a similar type
// Made/Build with 3 Headder elements, consists of
// 'Address pointing to underlying data-Structure of a Kind'
// 'Length of the Slice'
// 'Length of the underlaying Data-Structure'

// We would end-up having index out of range errors if an element is added at a random index
// best practiece is to use append, 'Go' will identify the existing capacity and double-down
// There is no direct way to delete items of a slice but to re-append as below.
//If i is the index of the element to be removed, then the format of this process would look like the following:
// slice = append(slice[:i], slice[i+1:]...)

// MultiDimensionalSlice demonstrates ways to insert a slice into a slice.
func MultiDimensionalSlice() {

	var student1 = make([]string, 3)
	var student2 = make([]string, 3)
	var student3 = make([]string, 3)

	student1[0] = "K"
	student1[1] = "N"
	student1[2] = "DOB"
	student2[0] = "S"
	student2[1] = "P"
	student2[2] = "DOB"
	student3[0] = "Randy"
	student3[1] = "Paush"
	student3[2] = "21-DEV-1971"

	// Declaring a multi dimensional slice in short-hand way <- casual approach
	// Takes slice of string as input but not pointing to an underlying array.
	var StudentYear2018 [][]string
	StudentYear2018 = append(StudentYear2018, student2)

	// A NIL slice will be created, which does not point to anything yet , till it gets initialized
	// at this point the underlaying array is not yet created.
	// No underlying array is created, NIL pointing
	// Most Prefered and Idiomatic way to make a Slice
	// reserved spots are already pointed in the memory
	// an underlying array is already created with the capacity size of 20
	StudentYear2020 := make([][]string, 0)
	// Throws an errow if initialized as below
	// StudentYear2020[0][0] = student3[0]
	// StudentYear2020[0][1] = "Paush"
	// StudentYear2020[0][1] = "N"
	// StudentYear2020[0][2] = "DOB"
	// elements can be directly added
	StudentYear2020 = append(StudentYear2020, student3)
	fmt.Printf("%v", StudentYear2020[0][0])

	// Increment a value in the slice
	ExampleSlice := make([]int, 10, 20)
	ExampleSlice[0] = 1
	//ExampleSlice[0] = ExampleSlice[0] +1
	//ExampleSlice[0] += n // <- 'n' being any similar nuber type
	ExampleSlice[0]++ // <- Would increment the value inside the slice idiomatic way

}

func main() {
	for i = 1; i <= 100; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Length:%d\tCapacity:%d\tValue:%v\n", len(mySlice), cap(mySlice), i)
	}

	//To delete the element at the position of 56
	mySlice = append(mySlice[:56], mySlice[57:]...)
	fmt.Printf("Length:%d\tCapacity:%d\nCompleteSliceValues:\t%v\n", len(mySlice), cap(mySlice), mySlice)
}
