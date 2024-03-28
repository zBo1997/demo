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
//
//	@param *option s
type Option func(*option)

// getOption 获取对象
//
//	@return *option
func getOption() *option {
	return cache.Get().(*option)
}

// releaseOption 重置对象
//
//	@param opt
func releaseOption(opt *option) {
	opt.reset()
	cache.Put(opt)
}

// WithSex WithSex
//
//	@param sex
//	@return Option
func WithSex(height int) Option {
	return func(o *option) {
		o.height = height
	}
}

// WithHobby set up Hobby
func WithHobby(hobby string) Option {
	return func(opt *option) {
		opt.hobby = hobby
	}
}

// WitWhWeight
//
//	@param weight
//	@return Option
func WitWhWeight(weight int) Option {
	return func(o *option) {
		o.weight = weight
	}
}

// WithAge WithAge
//
//	@param age 年龄
//	@return Option
func WithAge(age int) Option {
	return func(o *option) {
		o.age = age
	}
}

func findFriend(postion string, option ...Option) (string, error) {
	friend := fmt.Sprintf("从 %s 找朋友 \n", postion)
	//从缓存池获取对象
	opt := getOption()

	defer func() {
		releaseOption(opt)
	}()
	for _, o := range option {
		o(opt)
	}
	if opt.sex == 1 {
		sex := "性别：女性"
		friend += fmt.Sprintf("%s\n", sex)
	}
	if opt.sex == 2 {
		sex := "性别：男性"
		friend += fmt.Sprintf("%s\n", sex)
	}

	if opt.age != 0 {
		age := fmt.Sprintf("年龄：%d岁", opt.age)
		friend += fmt.Sprintf("%s\n", age)
	}

	if opt.height != 0 {
		height := fmt.Sprintf("身高：%dcm", opt.height)
		friend += fmt.Sprintf("%s\n", height)
	}

	if opt.weight != 0 {
		weight := fmt.Sprintf("体重：%dkg", opt.weight)
		friend += fmt.Sprintf("%s\n", weight)
	}

	if opt.hobby != "" {
		hobby := fmt.Sprintf("爱好：%s", opt.hobby)
		friend += fmt.Sprintf("%s\n", hobby)
	}
	return friend, nil
}

func main() {

	friend, _ := findFriend("附近的人",
		WithAge(1),
		WitWhWeight(180),
		WithHobby("滑雪"),
		WithSex(1))

	fmt.Println(friend)
}
