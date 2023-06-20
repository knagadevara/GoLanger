package main

import "fmt"

func PutDataOnChannel(nums ...int64) chan int64 {
	var myChn = make(chan int64)
	go func() {
		for _, num := range nums {
			myChn <- num
		}
		close(myChn)
	}()
	return myChn
}

func PrintItOut(ListofNums chan int64) {
	for chn1 := range ListofNums {
		fmt.Printf("Result: %d\n", chn1)
	}
}

func main() {
	chn := PutDataOnChannel(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// for num := range chn {
	// 	fmt.Printf("%d", num)
	// }
	PrintItOut(chn)
}
