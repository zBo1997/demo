package main

import (
	"fmt"
)

func main() {
	// recover 只能在defer中使用 而且defer只能捕获到当前协程的panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	go func() {

		panic("panic")
	}()

	select {}
}

// 正确使用的方式
func correct() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic:", err)
			}
		}()
		panic("panic in goroutine")
	}()
}
