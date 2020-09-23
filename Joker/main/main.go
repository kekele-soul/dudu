package main

import (
	"Joker/HandleRequest"
	"fmt"
	"net/http"
)


func main() {
	//fmt.Printf("\t\t\t\t欢迎进入")

	http.Handle("/static/img",http.StripPrefix("/static/img",http.FileServer(http.Dir(".HTML/"))))
	http.HandleFunc("/random",HandleRequest.Visit)
	http.HandleFunc("/time", HandleRequest.Xfei)
	http.HandleFunc("/login", HandleRequest.Xfei)
	http.HandleFunc("/new",HandleRequest.Shy)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println(err.Error())

	}

}

