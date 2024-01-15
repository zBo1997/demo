package main

import (
	"log"
	"net/http"
)

func main() {
	//配置路由
	/**
	 * 实现一个简单的Http 服务器请求
	 */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
