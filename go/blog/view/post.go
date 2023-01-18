package view

import (
	"blog/models"
	"blog/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlHandler) Post(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, _ := strconv.Atoi(pIdStr)

	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		models.Pages.Detail.WriteData(w, errors.New("查询文章详情失败"))
	} else {

		models.Pages.Detail.WriteData(w, postRes)
	}

}
