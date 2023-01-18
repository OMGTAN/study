package dao

import (
	"blog/common"
	"blog/models"
)

func GetPostTotal() int {
	rows := DB.QueryRow("select count(1) from blog_post ")
	var total int
	rows.Scan(&total)
	return total
}

func GetPostTotalBycategoryId(categoryId int) int {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ? ", categoryId)
	var total int
	rows.Scan(&total)
	return total
}

func GetPost(pageNo, pageSize int) []models.Post {
	pageNo = (pageNo - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", pageNo, pageNo+pageSize)
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

func GetPostByCategoryId(categeryId, pageNo, pageSize int) []models.Post {
	pageNo = (pageNo - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id=? limit ?,?", categeryId, pageNo, pageNo+pageSize)
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

func GetPostDetailById(pId int) (*models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid=?", pId)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	var post models.Post
	err = row.Scan(
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
	if err != nil {
		return nil, err
	}

	return &post, nil
}
