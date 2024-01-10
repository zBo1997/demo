package main

import (
	"fmt"
	"net"
)

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

/**
 * @Description:获取连接
 * @param 传入连接
 */
func handleConn(conn net.Conn) {
	// 函数调用完毕，自动关闭conn
	defer conn.Close()
	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect successful")
	// 循环读取客户端发送的数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
	}
}
