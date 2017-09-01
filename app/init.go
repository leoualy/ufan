package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
	"path"
)

type ConfDefault struct {
	Host string `host`
	Port int    `port`
	Url  string `url`
}

func getExecPath() string {
	curUser, err := user.Current()
	dir := curUser.HomeDir
	if nil == err {
		return dir
	}
	return "err"
}

func (c *ConfDefault) getConf() *ConfDefault {
	fmt.Println("当前用户:", getExecPath())
	// 组合配置文件路径
	dirCfg := path.Join(getExecPath(), "default.yaml")
	fileCfg, err := ioutil.ReadFile(dirCfg)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(fileCfg, c)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func main() {
	var c ConfDefault
	c.getConf()
	fmt.Println(c)
}
