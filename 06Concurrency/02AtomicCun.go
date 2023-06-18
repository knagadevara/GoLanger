package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var counter int32

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func incrementor(s string, till int) {
	for i := 0; i < till; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		{
			mutex.Lock()
			atomic.AddInt32(&counter, 1)
			fmt.Printf("%s: %d\tCounter: %d\n", s, i, counter)
			mutex.Unlock()
		}
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(2)
	go incrementor("NIM", 30)
	go incrementor("ROD", 30)
	waitGroup.Wait()
	fmt.Printf("Total: %d\n", counter)
}
