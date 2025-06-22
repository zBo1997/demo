// 每个go文件都需要归属于一个包
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := make([]int, 0, 3)
	b = append(b, a...) // 使用append函数将切片a的所有元素添加到切片b中
	a[0] = 100
	fmt.Println(a, b)
	fmt.Printf("%p, %p\n", a, b)

	//或者使用cpoy进行拷贝
	slice1 := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// 创建一个新的切片，长度和容量都与slice1相同 copy只能拷贝目标数组长度内的数据
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	//尝试引用数据类型copy ,当内部含有引用数据类型也可以拷贝
	slice3 := make([]Person, 2, 3)
	slice4 := []Person{
		{Name: "张三", Age: 18, bonus: Bonus{Name: "年终奖", Score: 1.2}},
		{Name: "李四", Age: 20, bonus: Bonus{Name: "年终奖", Score: 1.8}}}
	copy(slice3, slice4)
	fmt.Println(slice3, slice4)
}

type Person struct {
	Name  string
	Age   int
	bonus Bonus // 嵌套结构体
}

type Bonus struct {
	Name  string
	Score float32
}
