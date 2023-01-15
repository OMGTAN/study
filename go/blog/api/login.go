package api

import (
	"blog/models"
	"blog/service"
	"errors"
	"net/http"
)

func (*apiHdlStruct) Login(w http.ResponseWriter, r *http.Request) {
	loginRes, e := service.LoginService(r)
	if e != nil {
		models.Res.Fail(w, errors.New("查询用户失败"))
		return
	}
	models.Res.Success(w, loginRes)
}
