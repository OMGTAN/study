package api

import (
	"blog/models"
	"blog/service"
	"net/http"
)

func (*apiHdlStruct) SearchPost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	val := r.Form.Get("val")

	searchResp := service.SearchPost(val)

	models.Res.Success(w, searchResp)
}
