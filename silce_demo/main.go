// 每个go文件都需要归属于一个包
package main

import (
	"fmt"
)

func main() {
	var sli_1 []int //一个int 数组的切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)

	var sli_2 = [] int {}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_2),cap(sli_2),sli_2)

	var sli_3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_3),cap(sli_3),sli_3)

	for _, v := range sli_3 {
		fmt.Print(v)
		fmt.Print()
	}
	//切片sli_3截取 0~3 的位置
	sli_4 := sli_3[0:3]
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_4),cap(sli_4),sli_4)

	//创建一个切片，长度为3，容量为5
	test := make([]int32, 3, 5)

	for index,data  := range test {
		fmt.Println(index, data)
	}

	data := map[string]interface{}{}

	data["name"] = "张三"

	fmt.Println(data)
	
	for k, v := range data {
		fmt.Println(k, v)
	}

	//创建一个存放任何类型的切片
	all_type_slice := make([]interface{}, 10)
	all_type_slice = append(all_type_slice, 1)
	all_type_slice = append(all_type_slice, "hello")
	all_type_slice = append(all_type_slice, 1.2)
	all_type_slice = append(all_type_slice, true)

	//遍历切片
	for _, v := range all_type_slice {
		fmt.Println(v)
	}
}
