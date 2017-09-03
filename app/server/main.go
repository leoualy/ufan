package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"ufan/app/cfg"
)

type Account struct {
	Name string
	Age  int
}

func indexhtml(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Println(err)
	}
	var a = Account{}
	a.Name = `John`
	a.Age = 23
	tmpl.Execute(w, a)
}

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexhtml(w)
	})
	c := cfg.GetCfgDefault()
	portStr := strconv.Itoa(c.Port)
	hostStr := c.Host
	s := []string{hostStr, ":", portStr}
	addrStr := strings.Join(s, "")
	fmt.Println(addrStr)
	fmt.Println(c.TemplatePath)
	fmt.Println("服务运行中...")
	http.ListenAndServe(addrStr, nil)
}
