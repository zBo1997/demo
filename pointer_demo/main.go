package main

import (
	"fmt"
)

type student struct {
	id int
	name string
	age int
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
}
