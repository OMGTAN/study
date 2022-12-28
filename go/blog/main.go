package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type response struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func handleIndex(w http.ResponseWriter, request *http.Request) {
	res := &response{
		Title: "blog",
		Desc:  "desc",
	}
	jsonstr, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonstr)
}

func handleIndexHtml(w http.ResponseWriter, request *http.Request) {
	res := &response{
		Title: "你好 blog",
		Desc:  "首页测试",
	}
	t := template.New("index.html")
	path, _ := os.Getwd()
	t.ParseFiles(path + "/pc/index.html")
	t.Execute(w, res)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//处理'/'请求
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index.html", handleIndexHtml)

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
