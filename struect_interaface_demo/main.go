package main

import (
	"fmt"
)

type Study interface {
	// 1. 接口中的方法不能有方法体
	linsten(msg string) string

	speak(msg string) string

	read(msg string) string

	write(msg string) string
}

type study struct {
	Name string
}

// read implements Study.
func (*study) read(msg string) string {
	panic("unimplemented")
}

// write implements Study.
func (*study) write(msg string) string {
	panic("unimplemented")
}

func (*study) linsten(msg string) string {
	return "linsten: " + msg
}

func (*study) speak(msg string) string {
	return "speak: " + msg
}

func new(name string) Study {
	return &study{Name: name}
}

func main() {
	study := new("momo")
	fmt.Println(study.speak("hello"))
}
