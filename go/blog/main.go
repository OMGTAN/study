package main

import (
	"blog/config"
	"blog/models"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func date(format string) string {
	return time.Now().Format(format)
}

func handleIndexHtml(w http.ResponseWriter, request *http.Request) {
	t := template.New("index.html")
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date})
	path, _ := os.Getwd()
	index := path + "/template/index.html"
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	postList := path + "/template/layout/post-list.html"
	_, err := t.ParseFiles(index, home, header, footer, pagination, personal, postList)
	if err != nil {
		log.Println("解析模板失败：", err)
	}
	//返回数据
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var homeResp = &models.HomeResp{
		Viewer:    config.Conf.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	t.Execute(w, homeResp)
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
