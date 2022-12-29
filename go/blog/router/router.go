package router

import (
	"blog/view"
	"net/http"
)

func Route() {

	//处理'/'请求
	http.HandleFunc("/", view.HtmlHdl.Index)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
