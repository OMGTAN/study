package dao

import (
	"blog/common"
	"blog/models"
)

func GetUser(username, passwd string) *models.LoginRes {
	rows := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ?", username, passwd)

	var user models.User
	err := rows.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	common.PrintErr(err)
	if err != nil {
		return nil
	}
	token := "11"
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}

	return lr
}
