package main

import (
	"fmt"
	"time"
	"unicode"
)

var meterUp int = 0
var waitOnMe func(time.Duration)
var capMyRune func([]rune) []rune

func counter(tillCount int, waitTill time.Duration) {
	for i := 0; i <= tillCount+1; i++ {
		waitOnMe(waitTill)
		fmt.Printf("%v\n", i)
	}
}

func main() {

	waitOnMe = func(waitTill time.Duration) {
		time.Sleep(waitTill * time.Millisecond)
		meterUp += 1
	}

	capMyRune = func(listOfRunes []rune) []rune {
		capRuneArray := make([]rune, len(listOfRunes))
		for ix, char := range listOfRunes {
			capRuneArray[ix] = unicode.ToUpper(char)
			fmt.Printf("RuneToUpper: %c\n", capRuneArray[ix])
		}
		return capRuneArray
	}
	testRunes0 := []rune{'a', 'b', 'c'}
	testRunes1 := []rune{'d', 'e', 'f'}
	testRunes2 := []rune{'g', 'h', 'i'}
	testRunes3 := []rune{'j', 'k', 'l'}
	testRunes4 := []rune{'m', 'n', 'o'}
	testRunes5 := []rune{'p', 'q', 'r'}
	testRunes6 := []rune{'s', 't', 'u'}
	testRunes7 := []rune{'v', 'w', 'x', 'y', 'z'}
	go counter(10, 1000)
	go capMyRune(testRunes0)
	go counter(10, 100)
	go capMyRune(testRunes1)
	go capMyRune(testRunes2)
	go counter(10, 1000)
	go capMyRune(testRunes3)
	go capMyRune(testRunes4)
	go capMyRune(testRunes5)
	go counter(10, 1000)
	go capMyRune(testRunes6)
	go counter(10, 500)
	go capMyRune(testRunes7)
	fmt.Printf("Wait till completion\nMeterUp: %v\n", meterUp)
	time.Sleep(1100 * time.Millisecond)
	fmt.Printf("The End hogaya\nMeterUp: %v\n", meterUp)
}
