package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"ufan/app/cfg"
)

func Start() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ufan index")
	})
	c := cfg.GetCfgDefault()
	portStr := strconv.Itoa(c.Post)
	hostStr := c.Host
	s := []string{hostStr, ":", portStr}
	addrStr := strings.Join(s, "")
	fmt.Println(addrStr)
	fmt.Println("服务运行中...")
	http.ListenAndServe(addrStr, nil)
}
