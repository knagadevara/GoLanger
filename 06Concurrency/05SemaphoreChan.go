package main

import (
	"fmt"
)

var chn = make(chan int)
var done = make(chan bool)

func main() {
	go func() {
		for i := 1; i <= 10; i++ {
			chn <- i
		}
		done <- true
	}()

	go func() {
		for i := 51; i <= 60; i++ {
			chn <- i
		}
		done <- true
	}()

	// Wrongway
	// <-done
	// <-done
	// close(chn)

	// Rightway
	go func() {
		<-done
		<-done
		close(chn)
	}()

	// this loop will immediately start recieving the contents from channel buffer
	for n := range chn {
		fmt.Printf("%v\n", n)
	}

}
