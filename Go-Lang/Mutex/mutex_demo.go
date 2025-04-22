package main

import (
	"fmt"
	"sync"
)

var (
	sum  int
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 10; i++ {
		lock.Lock()
		sum += i
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(sum)
}
