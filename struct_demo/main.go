package main

import "fmt"

type student struct {
	id int
	name string
	age int
	sex string
	score int
	addr string
}

/** 
 * @Description:
 * @param student
 * @return student
 */
func test_struct(student student) student {
	student.id = 666
	student.addr = ""
	return student
}

func array_struct(stu[] student){
	if stu == nil {
		return
	}
	stu[2] = student{2, "李四", 18, "男", 100, "北京市海淀区"};
}


func main() {
	s := test_struct(student{1, "张三", 18, "男", 100, "北京市海淀区"})

	stu := []student{{2, "李四", 18, "男", 100, "北京市海淀区"}, {3, "王五", 18, "男", 100, "北京市海淀区"}, {4, "赵六", 18, "男", 100, "北京市海淀区"}}

	fmt.Println(s)
	fmt.Println(stu)
	array_struct(stu)
	fmt.Println(stu)
}
