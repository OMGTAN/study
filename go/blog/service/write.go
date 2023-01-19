package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
)

func GetWriting() *models.WritingRes {
	categorys := dao.GetAllCategory()
	writeRes := &models.WritingRes{
		Title:     config.Conf.Viewer.Title,
		CdnURL:    config.Conf.System.CdnURL,
		Categorys: categorys,
	}
	return writeRes
}
