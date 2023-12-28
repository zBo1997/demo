package main

import "fmt"

func main() {
	//声明数组
	var array1 = [5]int{1, 2, 3, 4, 5}
	array2 := [...]int{1, 2, 3, 4, 5}
	var arraydemo [10]int

	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(arraydemo))


	//声明一个切片 截取数组array1的第2个到第4个元素
	slice1 := array1[1:4]
	fmt.Println(slice1)
	//输出切片slice的长度
	fmt.Println(len(slice1))

	slice1 = append(slice1, array1[4:]...)
	fmt.Println(slice1)

	//修改slice1切片中第3个元素为0
	slice1[2] = 0
	//打印这个切片内容
	fmt.Println(slice1)

	str := "你好！我叫朱博"
	//把str转换成切片
	slice2 := []byte(str)
	//吧siice2中的"朱博"替换成"朱博博"
	//打印slice2
	fmt.Println(slice2)
	//以字符串形式打印slice2
	fmt.Println(string(slice2))

}
