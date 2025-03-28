package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage()
	fmt.Println("[main]Hello World!")
	//这里需要睡眠一段时间，否则主线程退出，goroutine无法执行
	time.Sleep(1 * time.Second)
}

func printMessage() {
	fmt.Println("Hello World!")
}
