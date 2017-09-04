package server

import (
	"fmt"
	"strconv"
	"time"
	"ufan/app/cfg"
	"ufan/app/db"
)

func checkDB(c cfg.CfgDefault) error {
	fmt.Println("数据库检测...")
	fmt.Println("数据库类型:", c.Driver)
	fmt.Println("数据库名:", c.Dbname)
	fmt.Println("数据库用户名:", c.User)
	dataSource := "user=" + c.User + " dbname=" + c.Dbname + " sslmode=disable"
	fmt.Println("数据源名称:", dataSource)
	err := db.Open(c.Driver, dataSource)
	if err != nil {
		fmt.Println("数据库初始化失败,错误消息:", err.Error())
		return err
	}
	fmt.Println("数据库服务正常!")
	return nil
}

var SrvAddrs string

func Init() error {
	fmt.Println("初始化...")
	c := cfg.GetCfgDefault()
	err := checkDB(c)
	SrvAddrs = c.Host + ":" + strconv.Itoa(c.Port)
	fmt.Println("初始化完成...")
	fmt.Println("当前时间:", time.Now())
	fmt.Println("服务地址:http://", SrvAddrs)
	return err
}
