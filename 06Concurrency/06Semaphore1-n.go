package main

import (
	"fmt"
)

const ChannelPollers int = 100000

var chn = make(chan int, 10)
var done = make(chan bool)

func main() {
	go func() {
		for i := 1; i <= 1000000; i++ {
			chn <- i
		}
		close(chn)
	}()

	for i := 0; i <= ChannelPollers; i++ {
		go func() {
			for n := range chn {
				fmt.Printf("%v\n", n)
			}
			done <- true
		}()
	}

	for i := 0; i <= ChannelPollers; i++ {
		<-done
	}
}
