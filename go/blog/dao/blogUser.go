package dao

import "blog/common"

func GetUserNameById(userId int) string {
	rows, err := DB.Query("select user_name from blog_user where uid = ?", userId)
	common.PrintErr(err)

	var userName string
	errr := rows.Scan(&userName)
	common.PrintErr(errr)
	return userName
}
