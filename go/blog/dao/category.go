package dao

import (
	"blog/common"
	"blog/models"
)

func GetAllCategory() []models.Category {
	rows, err := DB.Query("select * from blog_category")
	common.PrintErr(err)

	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		common.PrintErr(err)
		categorys = append(categorys, category)
	}
	return categorys
}
