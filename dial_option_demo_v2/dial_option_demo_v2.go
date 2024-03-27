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



// reset 重置对象参数
//
//	@receiver option
func (option *option) reset() {
	option.sex = 0
	option.age = 0
	option.height = 0
	option.weight = 0
	option.hobby = ""
}

// Option 声明一个类型是一个函数时接口，传递option的
//  @param *option s
type Option func(*option)

// getOption 获取对象
//  @return *option 
func getOption() *option{
	return cache.Get().(*option)
}

// releaseOption 重置对象
//  @param opt 
func releaseOption(opt *option){
	opt.reset()
	cache.Put(opt)
}

func WithSex(sex int)Option {
	return func(o *option) {
		o.sex = sex
	}
}

func main() {

	cache.Put(&option{name: "Tom"})
	fmt.Println(cache.Get())
}
