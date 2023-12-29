package main

import "fmt"

type stu struct {
	id   int
	name string
}

func main() {

	var s *stu
	s = new(stu)
	fmt.Println(s)
	println(s)

	*s = stu{1, "tom"}
	fmt.Println(*s)

}
