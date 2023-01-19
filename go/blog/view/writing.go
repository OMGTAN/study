package view

import (
	"blog/models"
	"blog/service"
	"net/http"
)

func (*HtmlHandler) Writing(w http.ResponseWriter, r *http.Request) {

	writing := models.Pages.Writing

	writeRes := service.GetWriting()

	writing.WriteData(w, writeRes)

}
