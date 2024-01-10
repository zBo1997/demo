package main

import (
	"fmt"
)

type student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {

	student := student{1, "张三", 18, "北京市海淀区"}

	var test_nil *float32
	fmt.Println(test_nil)
	student2 := &student

	a := 10
	b := 10

	p := &a

	*p = 666

	fmt.Println(*p)

	fmt.Println(p)
	fmt.Println(b)
	fmt.Println(&student2)

	fmt.Println("*****************************************************")

	//多级指针 p1 变量存放了 p 的地址
	var p1 **int = &p

	fmt.Println(p1)

	//p2 变量存放了 p1 的地址
	var p2 ***int = &p1

	fmt.Println(p2)

	***p2 = 777
	fmt.Println(&p2)

	fmt.Println(p2)
	fmt.Println(&p1)
}
