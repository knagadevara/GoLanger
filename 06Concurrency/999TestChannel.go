package main

import (
	"fmt"
	"runtime"
)

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
	PrintItOut(chn)
	fmt.Printf("%v\n", runtime.NumCPU())
}
