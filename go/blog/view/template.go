package view

import (
	"blog/config"
	"blog/models"
	"html/template"
	"log"
	"sync"
	"time"
)

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func date(format string) string {
	return time.Now().Format(format)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func LoadTemplate() {
	path := config.Conf.System.CurrentDir

	for _, value := range models.TemplateArr {

		t := template.New(value + ".html")
		t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date, "dateDay": DateDay})
		index := path + "/template/" + value + ".html"
		home := path + "/template/home.html"
		header := path + "/template/layout/header.html"
		footer := path + "/template/layout/footer.html"
		pagination := path + "/template/layout/pagination.html"
		personal := path + "/template/layout/personal.html"
		postList := path + "/template/layout/post-list.html"
		_, err := t.ParseFiles(index, home, header, footer, pagination, personal, postList)
		if err != nil {
			log.Println("解析模板"+value+"失败：", err)
		}

		var tm = models.TemplateModel{}
		tm.Template = t
		models.TempList = append(models.TempList, tm)
	}

	models.Pages.Category = models.TempList[0]
	models.Pages.Custom = models.TempList[1]
	models.Pages.Detail = models.TempList[2]
	models.Pages.Home = models.TempList[3]
	models.Pages.Index = models.TempList[4]
	models.Pages.Login = models.TempList[5]
	models.Pages.Pigeonhole = models.TempList[6]
	models.Pages.Writing = models.TempList[7]
}

func init() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		LoadTemplate()
		w.Done()
	}()
	w.Wait()
}
