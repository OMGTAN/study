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

	//处理'/p'请求,文章详情
	http.HandleFunc("/p/", view.HtmlHdl.Post)

	//处理写文章详情
	http.HandleFunc("/writing", view.HtmlHdl.Writing)

	//处理发布文章请求
	http.HandleFunc("/pigeonhole", view.HtmlHdl.Pigeonhole)

	//处理登录请求
	http.HandleFunc("/api/v1/login", api.ApiHdl.Login)

	//处理发布文章请求
	http.HandleFunc("/api/v1/post", api.ApiHdl.SaveOrUpdatePost)

	//处理发布文章请求
	http.HandleFunc("/api/v1/post/search", api.ApiHdl.SearchPost)

	//查询文章信息
	http.HandleFunc("/api/v1/post/", api.ApiHdl.GetPostDetail)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	//处理图片请求异常
	http.Handle("/p/img/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
