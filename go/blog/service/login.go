package service

import (
	"blog/common"
	"blog/dao"
	"blog/models"
	"errors"
	"net/http"
)

func LoginService(r *http.Request) (*models.LoginRes, error) {
	m := common.GetRequestJsonParam(r)
	username := m["username"].(string)
	passwd := m["passwd"].(string)
	loginRes := dao.GetUser(username, passwd)
	if loginRes == nil {
		e := errors.New("查询用户失败")
		return nil, e
	}
	return loginRes, nil
}
