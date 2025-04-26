// 每个go文件都需要归属于一个包
package main

import "fmt"

func main() {
	var sli_1 []int //一个int 数组的切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)

	var sli_2 = [] int {}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_2),cap(sli_2),sli_2)

	var sli_3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_3),cap(sli_3),sli_3)

	//创建一个unint
	var a uint32 = 1;
	var b uint32 = 0;
	fmt.Println("result=%d",b-a)
	fmt.Println("result=%d",a-b)
}
