package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type response struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func handleIndexHtml(w http.ResponseWriter, request *http.Request) {
	res := &response{
		Title: "你好 blog",
		Desc:  "首页测试",
	}
	t := template.New("index.html")
	path, _ := os.Getwd()
	log.Println(path)
	index := path + "/template/index.html"
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	postList := path + "/template/layout/post-list.html"
	t.ParseFiles(index, home, header, footer, pagination, personal, postList)
	t.Execute(w, res)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//处理'/'请求
	http.HandleFunc("/", handleIndexHtml)

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
