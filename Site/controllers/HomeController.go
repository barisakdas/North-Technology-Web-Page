package controllers

import (
	. "NorthTechWebPage/Log"
	. "NorthTechWebPage/Site/helpers"
	. "NorthTechWebPage/Site/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

type HomeController struct {

}

func (home HomeController) Index(w http.ResponseWriter, r *http.Request,params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory":func(categoryId int) string {
			return Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("home/")...)


 	//view, err := template.ParseFiles(Include("home/")...)
	if err != nil {
		LogJson("Site","Error","Homepage","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data:=make(map[string]interface{})
	data["about"] = About{}.GetAll()[0]
	data["services"] = Services{}.GetAll()[0:3]
	data["articles"] = Article{}.GetAll()[0:4]
	data["features"] = Feature{}.GetAll()[0:4]
	data["menus"] = Menu{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}
