package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)

	ch1 <- 0
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	//close(ch1)

	fmt.Printf("len(ch1): %d, cap(ch1): %d\n", len(ch1), cap(ch1))
	//使用goroutine读取ch1
	//依次取出ch1中的值 并且每次都判断ch1是否已经关闭 如果没有关闭就继续取值 如果已经关闭就退出循环
	go func() {
		for {
			v, ok := <-ch1
			if !ok {
				fmt.Println("ch1 is closed")
				break
			}
			fmt.Printf("v: %d, ok: %t\n", v, ok)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end")

}
