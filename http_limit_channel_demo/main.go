package main

import (
	"fmt"
	"net"
)

var sem = make(chan struct{}, 2)

func handleConn(c net.Conn) {
	sem <- struct{}{}
	defer func() { <-sem }()
	//do something
	fmt.Print("获取连接.....")
	fmt.Print("开始应答.....")
	c.Write([]byte("Hello World"))
}

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 函数调用完毕，自动关闭listener
	defer listener.Close()

	for {
		c, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			continue
		}
		//使用go 快速并发处理
		go handleConn(c)
	}
}
