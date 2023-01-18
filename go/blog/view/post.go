package view

import (
	"blog/models"
	"blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlHandler) Post(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	log.Println("1111111111111========", pIdStr)
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println("view post:", err.Error())
		models.Pages.Detail.WriteError(w, errors.New("pId转换失败"))
	}
	var postRes *models.PostRes
	postRes, err = service.GetPostDetail(pId)
	if err != nil {
		log.Println("view post:", err.Error())
		models.Pages.Detail.WriteError(w, errors.New("查询文章详情失败"))
	} else {

		models.Pages.Detail.WriteData(w, postRes)
		log.Println("postRes:", postRes.Viewer.Description)
	}

}
