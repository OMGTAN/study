package view

import (
	"blog/models"
	"blog/service"
	"net/http"
)

func (*HtmlHandler) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := models.Pages.Pigeonhole

	pigeonholeRes := service.GetPigeonhole()

	pigeonhole.WriteData(w, pigeonholeRes)
}
