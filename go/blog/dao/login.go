package dao

import (
	"blog/common"
	"blog/models"
	"log"
)

func GetUser(username, passwd string) *models.User {
	rows := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ?", username, passwd)

	var user models.User
	err := rows.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	common.PrintErr(err)
	log.Println(user.UpdateAt)
	if err != nil {
		return nil
	}

	return &user
}
