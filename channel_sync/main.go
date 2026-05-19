package main

import (
	"fmt"
	"time"
)

func task1(done chan bool) {
	fmt.Println("Task1 1 starter")
	time.Sleep(2 * time.Second)
	fmt.Println("Task1 1 end")
	done <- true
}

func task2(done chan bool) {
	fmt.Println("Task1 2 starter")
	time.Sleep(2 * time.Second)
	fmt.Println("Task1 2 end")
	done <- true
}

func main() {
	done := make(chan bool)

	go task1(done)
	go task2(done)

	//发送任务完成的信号
	<-done
	<-done

	fmt.Println("All task completed")
}
