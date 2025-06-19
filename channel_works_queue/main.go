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

func (this *IcuQueue) NewIcuQueue(ctx context.Context, size int) *IcuQueue {
	queue := &IcuQueue{}
	//初始化队列和队列中context上下文
	queue.ctx, queue.cancel = context.WithCancel(context.Background())
	queue.ch = make(chan func(), size)
	queue.waiter.Add(size)
	return queue
}

func (this *IcuQueue) Push(item func(), timeout int) error {
	timer := time.NewTimer(time.Duration(timeout) * time.Second) // 设置超时时间
	select {
	case this.ch <- item:
		return nil
	case <-timer.C:
		return fmt.Errorf("timeout after %d seconds, item not added to queue", timeout)
	default:
		return fmt.Errorf("queue is full, item not added to queue")
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
		default:
			fmt.Println("queue is empty, no item to pop")
			return // 如果队列为空，则直接返回
		}
	}
}

func (this *IcuQueue) Close() {
	//发出关闭信号
	this.cancel()
	//等待所有的任务完成
	this.waiter.Wait()
}
