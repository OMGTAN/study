package dao

import (
	"blog/common"
	"blog/models"
)

func GetPost(pageNo, pageSize int) []models.Post {
	pageNo = (pageNo - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit?,?", pageNo, pageSize)
	common.PrintErr(err)

	var postList []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		common.PrintErr(err)
		postList = append(postList, post)
	}

	return postList
}
