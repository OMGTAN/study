package view

import (
	"blog/config"
	"blog/models"
	"net/http"
)

func (*HtmlHandler) Index(w http.ResponseWriter, request *http.Request) {
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
	models.Pages.Index.WriteData(w, homeResp)
}
