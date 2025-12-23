package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println("Context value:", ctx.Value("key"))
}

func contextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Context before cancel:", ctx.Err())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine exiting:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	//发送取消信号
	cancel()
	time.Sleep(2 * time.Second) //等待goroutine退出
}

func contextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	fmt.Println("Context before deadline:", ctx.Err())

	//模拟任务超时
	//当任务超过2秒未完成，context会自动取消
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task timed out:", ctx.Err())
	}

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine exiting:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

}

func main() {
	contextWithValue()
	contextWithCancel()
	contextWithDeadline()
}
