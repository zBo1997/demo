package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	go func() {
		fmt.Print("goroutine 1\n")
	}()

	time.Sleep(1 * time.Second)

	fmt.Print("end\n")

	ch1 := make(chan int, 10)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4

	//close(ch1)

	fmt.Printf("len(ch1): %d, cap(ch1): %d\n", len(ch1), cap(ch1))
	//使用goroutine读取ch1
	go func() {
		//依次取出ch1中的值
		for v := range ch1 {
			fmt.Println(v)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end")

}
