package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func MakeSound(a Animal) string {
	return a.Speak()
}

func main() {
	var a Animal = Dog{}
	fmt.Println(MakeSound(a)) // Output: Woof!
}
