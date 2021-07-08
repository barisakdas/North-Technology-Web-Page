package controllers

import (
	. "NorthTechWebPage/Log"
	. "NorthTechWebPage/Site/helpers"
	. "NorthTechWebPage/Site/models"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type ServiceController struct {}

func (service ServiceController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(Include("service/")...)
	if err != nil {
		LogJson("Site","Error","Service","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data := make(map[string]interface{})
	data["services"]=Services{}.GetAll()
	data["features"] = Feature{}.GetAll()[0:4]
	data["menus"] = Menu{}.GetAll()
	view.ExecuteTemplate(w,"service",data)
}