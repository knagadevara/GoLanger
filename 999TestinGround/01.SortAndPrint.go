package main

import (
	"fmt"
)

var unSortedArray = make([]int, 0)
var TempArray = make([]int, 0)
var highestInt int

func GetLargest(gretNum []int, largest int) int {
	for _, numG := range gretNum {
		if numG > largest {
			largest = numG
		}
	}
	return largest
}

func main() {
	unSortedArray = []int{34, 11, 45, 91, 1, 87, 21, 43, 77}
	highestInt = GetLargest(unSortedArray, highestInt)
	TempArray = append(TempArray, highestInt)
	if len(TempArray) > 0 {
		fmt.Printf("%d\n", TempArray)
	}
}
