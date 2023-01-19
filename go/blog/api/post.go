package api

import (
	"blog/common"
	"blog/models"
	"blog/service"
	"blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*apiHdlStruct) SaveOrUpdatePost(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	_, claims, _ := utils.ParseToken(token)
	uid := claims.Uid
	params := common.GetRequestJsonParam(r)

	method := r.Method

	switch method {
	case http.MethodPost:
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		models.Res.Success(w, post)
	case http.MethodPut:
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		cId := params["categoryId"].(float64)
		categoryId := int(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		models.Res.Success(w, post)
	}

}

func (*apiHdlStruct) GetPostDetail(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, _ := strconv.Atoi(pidStr)
	post, _ := service.GetPostById(pid)

	models.Res.Success(w, post)
}
