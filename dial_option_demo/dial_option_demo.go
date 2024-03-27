package main

import "fmt"


// Add 
//  @param a 
//  @param args 
//  @return int 
func Add(a *int, args ...int) int {
	result := a
	for _, v := range args {
		*result += v
	}
	return *result
}


func main() {
	var p = 1

	fmt.Println(Add(&p, 2, 3))

	p = 1
	fmt.Println(p)
}
