package view

import (
	"blog/common"
	"blog/config"
	"blog/dao"
	"blog/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlHandler) Index(w http.ResponseWriter, request *http.Request) {
	var categorys = dao.GetAllCategory()

	request.ParseForm()
	pageStr := request.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	pageNo, err := strconv.Atoi(pageStr)
	common.PrintErr(err)
	pageSize := 10

	path := request.URL.Path
	slug := strings.TrimPrefix(path, "/")
	var posts []models.Post
	var total int
	if slug == "" {
		posts = dao.GetPost(pageNo, pageSize)
		total = dao.GetPostTotal()
	} else {
		posts = dao.GetPostBySlug(slug, pageNo, pageSize)
		total = dao.GetPostTotalBySlug(slug)

	}

	var postMoreList []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     DateDay(post.CreateAt),
			UpdateAt:     DateDay(post.UpdateAt),
		}
		postMoreList = append(postMoreList, postMore)
	}

	pageTotal := (total-1)/pageSize + 1
	var pages []int
	for i := 1; i < pageTotal+1; i++ {
		pages = append(pages, i)
	}
	homeResp := &models.HomeResp{
		Viewer:    config.Conf.Viewer,
		Categorys: categorys,
		Posts:     postMoreList,
		Total:     total,
		Page:      pageNo,
		Pages:     pages,
		PageEnd:   pageNo != pageTotal,
	}
	models.Pages.Index.WriteData(w, homeResp)
}
