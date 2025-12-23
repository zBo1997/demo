package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var (
	target      = flag.String("target", "http://theta.icu", "目标 URL")
	total       = flag.Int64("total", 10000, "总请求数")
	concurrency = flag.Int("concurrency", 200, "并发请求数")
)

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	var success, fail int64
	sem := make(chan struct{}, *concurrency)

	for i := int64(0); i < *total; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			resp, err := http.Get(*target)
			if err != nil {
				atomic.AddInt64(&fail, 1)
				return
			}
			resp.Body.Close()
			atomic.AddInt64(&success, 1)
		}()
	}

	// 状态打印
	go func() {
		for {
			fmt.Printf("[%s] 成功: %d, 失败: %d, 并发: %d\n",
				time.Now().Format("15:04:05"), atomic.LoadInt64(&success),
				atomic.LoadInt64(&fail), len(sem))
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Printf("测试完成，成功=%d, 失败=%d\n", success, fail)
}
