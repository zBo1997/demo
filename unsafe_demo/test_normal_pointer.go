package unsafedemo

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 10

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = i
	}
	fmt.Println(b)
	// [0 1 2 3 4 5 6 7 8 9]

	// 取slice的最后的一个元素
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + 9*unsafe.Sizeof(b[0]))
	// 等价于unsafe.Pointer(&b[9])
	fmt.Println(*(*int)(end))
	// 9

	fmt.Println(uintptr(1) << (2 & (8*8 - 1)))
	// 4
}
