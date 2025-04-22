package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum     int
	mutex   sync.Mutex
	wrmutex sync.RWMutex
	wg      sync.WaitGroup
)

func write() {
	wrmutex.Lock()
	sum += 1
	time.Sleep(time.Millisecond * 10)
	wrmutex.Unlock()
	wg.Done()
}

func read() {
	wrmutex.RLock()
	time.Sleep(time.Millisecond)
	wrmutex.RUnlock()
	wg.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
