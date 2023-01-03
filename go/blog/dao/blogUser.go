package dao

import "blog/common"

func GetUserNameById(userId int) string {
	rows := DB.QueryRow("select user_name from blog_user where uid = ?", userId)

	var userName string
	errr := rows.Scan(&userName)
	common.PrintErr(errr)
	return userName
}
