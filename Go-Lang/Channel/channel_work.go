package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d received job %d\n", id, job)
		results <- job * job
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d finished job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
}
