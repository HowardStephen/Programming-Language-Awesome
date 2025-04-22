package main

import (
	"fmt"
)

func Producer() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				ch <- i // 2, 4, 6, 8
			}
		}
		close(ch)
	}()

	return ch
}

func Consumer(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum // 20
}

func main() {
	ch := Producer()

	res := Consumer(ch)

	fmt.Println(res)
}
