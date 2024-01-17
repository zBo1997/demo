package main

import (
	"fmt"
	"io"
	"net"
)

var sem = make(chan struct{}, 2)

func handleConn(c net.Conn) {
	sem <- struct{}{}
	defer func() { <-sem }()
	//获取请求数据
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			//如果是读IO错误则关闭丽娜姐
			if err != io.EOF {
				fmt.Println("Failed to read from connection:", err)
			}
			return
		}
		//输出结果到控制台
		fmt.Println("Received:", string(buf[:n]))
		// 写入数据
		var input string
		//创建一个新的对象
		fmt.Scanln(&input)
		_, err = c.Write([]byte(input))
		if err != nil {
			fmt.Println("Failed to write to connection:", err)
			return
		}

	}
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
