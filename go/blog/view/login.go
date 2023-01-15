package view

import (
	"blog/config"
	"blog/models"
	"net/http"
)

func (*HtmlHandler) Login(w http.ResponseWriter, request *http.Request) {

	models.Pages.Login.WriteData(w, config.Conf.Viewer)

}
