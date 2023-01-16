package service

import (
	"blog/common"
	"blog/dao"
	"blog/models"
	"blog/utils"
	"errors"
	"log"
	"net/http"
)

func LoginService(r *http.Request) (*models.LoginRes, error) {
	m := common.GetRequestJsonParam(r)
	username := m["username"].(string)
	passwd := m["passwd"].(string)
	passwd = utils.Md5Crypt(passwd, "mszlu")
	log.Println("mima:" + passwd)
	user := dao.GetUser(username, passwd)

	if user == nil {
		e := errors.New("查询用户失败")
		return nil, e
	}
	token, err := utils.Award(&user.Uid)
	common.PrintErr(err)
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
