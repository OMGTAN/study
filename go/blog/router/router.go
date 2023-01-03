package router

import (
	"blog/view"
	"net/http"
)

func Route() {

	//处理'/'请求
	http.HandleFunc("/", view.HtmlHdl.Index)

	//处理'/c'请求,分类分页
	http.HandleFunc("/c/", view.HtmlHdl.Category)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
