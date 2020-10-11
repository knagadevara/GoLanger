package main

import (
	"fmt"
	"sort"
)

//TakeTwo will return two values 01
func TakeTwo(a int) (y int, x bool) {
	if a%2 == 0 {
		x = true
	} else {
		x = false

	}
	y = a / 2
	return
}

// SimpleTakeTwo same functionality but minimal code
func SimpleTakeTwo(a int) (int, bool) {

	return a / 2, a%2 == 0
}

var numList = make([]int, 20, 40)

//WhIsGreat is a sorting stat
func WhIsGreat(gretNum ...int) {

	sort.Ints(gretNum)
	lastNum := len(gretNum) - 1
	fmt.Printf("The greatest Number is: %v \n", gretNum[lastNum])

}

// SortingFunc the correct way.
func SortingFunc(gretNum ...int) int {
	var largest int
	for _, numG := range gretNum {

		if numG > largest {

			largest = numG
		}
	}

	return largest
}
