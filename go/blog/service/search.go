package service

import (
	"blog/dao"
	"blog/models"
)

func SearchPost(val string) []models.SearchResp {

	posts := dao.SearchPost(val)

	var result []models.SearchResp

	for _, post := range posts {

		result = append(result, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}

	return result
}
