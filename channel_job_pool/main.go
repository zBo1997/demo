package main

import (
	"fmt"
)

// 工作池模式，其中三个worker并行处理job
func main() {
	job := make(chan any, 100)
	result := make(chan any, 100)

	for i := 0; i < 3; i++ {
		go worker(i, job, result)
	}

	for i := 0; i < 5; i++ {
		job <- i
	}

	close(job)

	for i := 0; i < 6; i++ {
		<-result
	}
}

func worker(ud int, job chan any, result chan any) {
	for j := range job {
		fmt.Println("worker", ud, "processing job", j)
		result <- j
	}
}
