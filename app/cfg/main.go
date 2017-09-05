package cfg

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path"

	"gopkg.in/yaml.v2"
)

type CfgDefault struct {
	Host         string `host`
	Port         int    `port`
	TemplatePath string `templatepath`
	Driver       string `driver`
	User         string `user`
	Dbname       string `dbname`
	StaticDir    string `staticdir`
}

func GetHomeDir() string {
	curUser, err := user.Current()
	dir := curUser.HomeDir
	if nil == err {
		return dir
	}
	return err.Error()
}

func GetCfgDefault() CfgDefault {
	var c = CfgDefault{}
	dirCfg := path.Join(GetHomeDir(), "ufan-conf", "default.yaml")
	fileCfg, err := ioutil.ReadFile(dirCfg)
	if err != nil {
		fmt.Println(err)
		return c
	}
	err = yaml.Unmarshal(fileCfg, &c)
	if err != nil {
		fmt.Println(err)
		return c
	}
	return c
}
