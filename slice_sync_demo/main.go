// 每个go文件都需要归属于一个包
package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	slice := make([]int, 0, 10)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // 确保在协程结束时调用Done，减少WaitGroup计数
			lock.Lock()     // 锁定，防止并发写入
			slice = append(slice, i)
			lock.Unlock() // 解锁，允许其他协程写入
		}(i)
	}
	wg.Wait()
	fmt.Println(slice)
	fmt.Println(len(slice))
}
