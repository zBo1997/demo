package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type IcuQueue struct {
	ctx    context.Context
	cancel func()
	waiter sync.WaitGroup
	ch     chan func()
}

func (this *IcuQueue) NewIcuQueue(workers int,size int) *IcuQueue {
	queue := &IcuQueue{}
	//初始化队列和队列中context上下文
	queue.ctx, queue.cancel = context.WithCancel(context.Background())
	queue.ch = make(chan func(), size)
	for i := 0; i < workers; i++ {
		queue.waiter.Add(1)
		go queue.Pop() // 启动workers个消费者
	}
	return queue
}

func (this *IcuQueue) Push(item func(), timeout int) error {
	timer := time.NewTimer(time.Duration(timeout) * time.Second) // 设置超时时间
	select {
	case this.ch <- item:
		return nil
	case <-timer.C:
		return fmt.Errorf("timeout after %d seconds, item not added to queue", timeout)
	}
}

func (this *IcuQueue) Pop() {
	//执行完毕以后要设置waiter.Done()，表示当前的任务已经完成
	defer this.waiter.Done()
	for {
		select {
		case item := <-this.ch:
			item() // 执行队列中的函数
		case <-this.ctx.Done():
			fmt.Println("queue is closed, no item to pop")
			return
		}
	}
}

func (this *IcuQueue) Close() {
	//发出关闭信号
	this.cancel()
	//等待所有的任务完成
	this.waiter.Wait()
}


func main() {
   	workerNum := 5 // 用户可自定义
    queueSize := 10
	queue := new(IcuQueue).NewIcuQueue(workerNum,queueSize)

	// 向队列中添加任务
	for i := 0; i < 20; i++ {
		i := i
		err := queue.Push(func() {
			fmt.Printf("task %d is running\n", i)
			time.Sleep(2 * time.Second)
		}, 1)
		if err != nil {
			fmt.Println(err)
		}
	}
	queue.Close()
}