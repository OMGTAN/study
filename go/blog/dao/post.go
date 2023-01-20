package dao

import (
	"blog/common"
	"blog/models"
	"log"
)

func GetPostTotal() int {
	rows := DB.QueryRow("select count(1) from blog_post ")
	var total int
	rows.Scan(&total)
	return total
}

func GetPostTotalBySlug(slug string) int {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ? ", slug)
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

func GetPostBySlug(slug string, pageNo, pageSize int) []models.Post {
	pageNo = (pageNo - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?,?", slug, pageNo, pageNo+pageSize)
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

func GetAllPost() []models.Post {
	rows, err := DB.Query("select * from blog_post")
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

func SearchPost(val string) []models.Post {
	rows, err := DB.Query("select * from blog_post where title like ?", "%"+val+"%")
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
		log.Println("dao post 1111, pid: ", pId, " err msg: ", err.Error())
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
		log.Println("dao post 2222, pid: ", pId, " err msg: ", err.Error())
		return nil, err
	}

	return &post, nil
}

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}
