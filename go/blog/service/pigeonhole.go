package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
)

func GetPigeonhole() *models.PigeonholeRes {

	postList := dao.GetAllPost()
	var line = make(map[string][]models.Post)
	for _, post := range postList {
		att := post.CreateAt
		month := att.Format("2006-01")
		line[month] = append(line[month], post)
	}
	categorys := dao.GetAllCategory()
	pigeonholeRes := &models.PigeonholeRes{
		Viewer:       config.Conf.Viewer,
		SystemConfig: config.Conf.System,
		Categorys:    categorys,
		Lines:        line,
	}
	return pigeonholeRes
}
