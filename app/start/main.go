package main

import (
	"fmt"
	"ufan/app/server"
)

func main() {
	fmt.Println("开始启动服务...")
	fmt.Println("本地IP列表:")
	server.Start()
}
