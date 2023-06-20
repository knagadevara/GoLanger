package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseArgSloce(s []string) []uint64 {
	var listOfNum = make([]uint64, 0)
	var someuint uint64
	for _, num := range s[1:] {
		someuint, _ = strconv.ParseUint(num, 10, 64)
		listOfNum = append(listOfNum, someuint)
	}
	return listOfNum
}

func calculateSquare(num *uint64) *uint64 {
	somemul := (*num) * (*num)
	return &somemul
}
func calculateSquareCh(nums ...uint64) (sqChan *chan uint64) {
	var outChan = make(chan uint64)
	go func() {
		for _, num := range nums {
			outChan <- *calculateSquare(&num)
		}
		close(outChan)
	}()
	return &outChan
}

func calculateFactorial(num *uint64) *uint64 {
	var total uint64 = 1
	for index := *num; index > 0; index-- {
		total *= index
	}
	return &total
}
func calculateFactorialCh(nums ...uint64) (calFact *chan uint64) {
	var out = make(chan uint64)
	go func() {
		for _, v := range nums {
			out <- *calculateFactorial(&v)
		}
		close(out)
	}()
	return &out
}

func printItOut(un uint64) { fmt.Printf(":\t%d\n", un) }
func printItOutCh(ListofNums *chan uint64) {
	go func() {
		for chn1 := range *ListofNums {
			printItOut(chn1)
		}
	}()
}

func main() {
	printItOutCh(calculateFactorialCh(parseArgSloce(os.Args)...))
	printItOutCh(calculateSquareCh(parseArgSloce(os.Args)...))

}
