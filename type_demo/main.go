// 每个go文件都需要归属于一个包
package main

//下面用到了fmt的Println函数，因此需要导包
import (
	"fmt"
	"strconv"
)

func main() {

	a := 3
	yest := 20.2
	num := 32.2
	fmt.Println("hello,go" + strconv.Itoa(a))
	fmt.Println(yest)
	fmt.Println(num)

	//自动类型推断
	pi :=3.14
	fmt.Printf("%T\n",pi)
}
