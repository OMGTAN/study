package models

import (
	"html/template"
	"log"
	"net/http"
)

var TemplateArr []string = []string{"category", "custom", "detail", "home", "index", "login", "pigeonhole", "writing"}

var TempList []TemplateModel

type TemplateModel struct {
	*template.Template
}

type TemplateAll struct {
	Category   TemplateModel
	Custom     TemplateModel
	Detail     TemplateModel
	Home       TemplateModel
	Index      TemplateModel
	Login      TemplateModel
	Pigeonhole TemplateModel
	Writing    TemplateModel
}

var Pages = &TemplateAll{}

func (t *TemplateModel) WriteData(w http.ResponseWriter, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		log.Panicln("页面填充数据失败：", err)
	}
}

func (t *TemplateModel) WriteError(w http.ResponseWriter, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		log.Panicln("错误信息：", err)
	}
}
