package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"ufan/app/cfg"
)

var connString = ""

func open() int {
	var c cfg.CfgDefault
	if connString == "" {
		c = cfg.GetCfgDefault()
		connString = "user=" + c.User + " dbname=" + c.Dbname + " sslmode=disable"
	}

	fmt.Println(connString)
	_, err := sql.Open(c.Driver, connString)
	if err != nil {
		return -1
	}
	return 0
}

// 数据库连接测试
func TestDbConnection() {
	if open() == 0 {
		fmt.Println("数据库连接成功!")
		return
	}
	fmt.Println("数据库连接失败!")
}
