package main

import (
	"fmt"
	"sync"
)

var chn = make(chan int)
var waitGroup sync.WaitGroup
var mutex sync.Mutex

func main() {

	go func() {
		waitGroup.Add(1)
		for i := 0; i < 10; i++ {
			chn <- i
		}
		waitGroup.Done()
	}()

	go func() {
		waitGroup.Add(1)
		for i := 50; i < 60; i++ {
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
