package server

import (
	"fmt"
	"html/template"
	"net/http"
)

type Account struct {
	Name string
	Age  int
}

func indexhtml(w http.ResponseWriter) {
	tmpl, _ := template.ParseFiles("views/index.html")
	//if err != nil {
	//	fmt.Println(err)
	//}
	var a = Account{}
	a.Name = `John`
	a.Age = 23
	tmpl.Execute(w, a)
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
