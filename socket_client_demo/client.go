package main

import (
	"fmt"
	"net"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()
	fmt.Println("conn successful = ", conn)
	// 发送数据
	message := "hello world"
	conn.Write([]byte(message))

	//读取服务器回发的数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	string := string(buf[:n])
	fmt.Println("服务器回发：", string)
}
