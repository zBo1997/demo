// 每个go文件都需要归属于一个包
package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6] // 下标2到5（含2不含6）
	s2 := arr[3:5] // 下标5到6（含5不含7）
	fmt.Println(s1, s2)
}
