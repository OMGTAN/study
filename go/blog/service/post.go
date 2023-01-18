package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"html/template"
	"time"
)

func GetPostDetail(pId int) (*models.PostRes, error) {

	post, err := dao.GetPostDetailById(pId)

	if err != nil {
		return nil, err

	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     post.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:     post.UpdateAt.Format("2006-01-02 15:04:05"),
	}
	postRes := models.PostRes{
		Description:  "",
		Viewer:       config.Conf.Viewer,
		SystemConfig: config.Conf.System,
		Article:      postMore,
	}
	return &postRes, nil
}

func DateDay(time time.Time) {
	panic("unimplemented")
}
