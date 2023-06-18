package main

import (
	"fmt"
	"sync"
)

var chn = make(chan int)
var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)

	go func() {
		for i := 1; i <= 10; i++ {
			chn <- i
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 51; i <= 60; i++ {
			chn <- i
		}
		waitGroup.Done()
	}()

	go func() {
		waitGroup.Wait()
		close(chn)
	}()

	for n := range chn {
		fmt.Printf("%v\n", n)
	}

}
