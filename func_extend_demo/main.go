package main

type persion struct {
	name string
	id   int
	sex  string
	age  int
}

type student struct {
	persion

	score int
	add   string
}

func (p *persion) sayHello() {
	println("Hello, my name is", p.name)
}
//
func main() {
	//创建一个student对象
	var stu student = student{persion{"张三", 1, "男", 18}, 100, "北京市"}
	//调用persion的方法
	stu.sayHello()
}



