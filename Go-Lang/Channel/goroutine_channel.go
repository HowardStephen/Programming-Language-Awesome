package main

import "fmt"

func Produce(pro chan<- int) {
	for i := 0; i < 10; i++ {
		pro <- i
	}
	close(pro)
}

func Receive(rec chan<- int, pro <-chan int) {
	//for {
	//	x, ok := <- pro
	//	if !ok {
	//		break
	//	}
	//  rec <- x * x
	//	fmt.Println(x)
	//}
	// close(rec)
	for v := range pro {
		rec <- v * v
	}
	close(rec)
}

func main() {
	// 定义 + 初始化
	var pro chan int
	var rec chan int

	pro = make(chan int, 10)
	rec = make(chan int, 10)

	// pro := make(chan int, 10)
	// rec := make(chan int, 10)

	go Produce(pro)
	go Receive(rec, pro)

	for v := range rec {
		fmt.Println(v)
	}
}
