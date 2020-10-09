package main

import "fmt"

func main() {

	x, y := TakeTwo(5)
	fmt.Printf("(%v,%v)\n", y, x)

	// Doing the same with function expression

	j := func(a int) (y int, x bool) {
		if a%2 == 0 {
			x = true
		} else {
			x = false

		}
		y = a / 2
		return
	}

	a, b := j(10)
	fmt.Printf("(%v,%v)\n", a, b)
	numList = []int{12, 1, 6, 2, 98, 4, 56, 32}
	WhIsGreat(numList...)
	grtNum := SortingFunc(numList...)
	fmt.Println(grtNum)

}
