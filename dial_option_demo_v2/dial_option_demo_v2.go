package main

import (
	"fmt"
	"sync"
)

// option
type option struct {
	name   string
	sex    int
	age    int
	height int
	weight int
	hobby  string
}

// var
// 声明一个池子，池子中存放option对象
//
//	@param cache 对象池 内部是一个链表 线程安全的
//	@param 里面赋值了一个默认"option"对象
var (
	cache = &sync.Pool{
		New: func() interface{} {
			return &option{name: "Jerry"}
		},
	}
)

func main() {

	cache.Put(&option{name: "Tom"})
	fmt.Println(cache.Get())
	fmt.Println(cache.Get())
}
