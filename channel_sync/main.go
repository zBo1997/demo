package main

import (
	"fmt"
	"time"
)

// 此demochannel之间的同步
func main() {
	done := make(chan bool)

	go task1(done)
	go task2(done)

	<-done
	<-done
	<-done

	fmt.Println("All task completed")
}

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
