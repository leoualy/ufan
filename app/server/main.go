package server

import (
	"fmt"
	"net/http"
)

type Account struct {
	Name string
	Age  int
}

func indexhtml(w http.ResponseWriter) {
	var a = Account{}
	a.Name = `John`
	a.Age = 23
	View(w, "Views/index.html", a)
}

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexhtml(w)
	})
	fmt.Println("启动Ufan...")
	err := Init()
	if err != nil {
		fmt.Println("Ufan启动失败,错误消息", err)
		return
	}

	err = http.ListenAndServe(SrvAddrs, nil)
	if err != nil {
		fmt.Println("Ufan启动失败,错误消息", err)
		return
	}
	fmt.Println("Ufan启动成功!")
}
