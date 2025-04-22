package main

import "fmt"

func f(ch chan int) {
	for {
		v, o := <-ch
		if !o {
			fmt.Println("channel closed")
			break
		}
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	f(ch)
}
