package main

import (
	"fmt"
	"net"
	"ufan/app/cfg"
)

func showAddress() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("获取IP错误:", err)
	}

	for i, addr := range addrs {
		fmt.Println(i, addr)
	}
}

func main() {
	fmt.Println("开始启动服务...")
	fmt.Println("当前用户目录:", cfg.GetHomeDir())
	fmt.Println("开始加载配置...")
	fmt.Println("本地IP列表:")
	showAddress()
	c := cfg.GetCfgDefault()
	fmt.Println(c.Host)
}
