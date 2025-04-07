package main

import (
	"fmt"
	"sync"
	"time"
)

type Count struct {
	mu    sync.Mutex
	value int32
}

func (c *Count) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Count) GetValue() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {

	count := Count{
		value: 0,
		mu:    sync.Mutex{},
	}

	go count.Add()
	fmt.Println("Count value:", count.GetValue())

	time.Sleep(1 * time.Second) // 等待一段时间，确保所有 goroutine 完成
}
