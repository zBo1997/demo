package main

import "fmt"

type students struct {
	id   int
	name string
	age  int
	sex  string
}

type Int int

// func (方法接收者 数据类型) 方法名(方法参数列表) 返回值列表{代码体}
func (a Int) add(b Int) Int {
	return a + b
}

// 把名字当前名字设置为空
func emptyName(s *students) {
	s.name = ""
}

func main() {
	var a Int = 1
	i := a.add(2)
	fmt.Println(i)

	stu := &students{1, "momo", 18, "男"}
	fmt.Println(stu)
	emptyName(stu)
	fmt.Println(stu)
}
