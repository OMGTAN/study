package router

import (
	"blog/api"
	"blog/view"
	"net/http"
)

func Route() {

	//处理'/'请求
	http.HandleFunc("/", view.HtmlHdl.Index)

	//处理登录页面请求
	http.HandleFunc("/login", view.HtmlHdl.Login)

	//处理'/c'请求,分类分页
	http.HandleFunc("/c/", view.HtmlHdl.Category)

	//处理'/c'请求,分类分页
	http.HandleFunc("/api/v1/login", api.ApiHdl.Login)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
