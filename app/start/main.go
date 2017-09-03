package main

import (
	"fmt"
	"ufan/app/db"
	"ufan/app/server"
)

func main() {
	fmt.Println("开始启动服务...")
	db.TestDbConnection()
	server.Start()
}
