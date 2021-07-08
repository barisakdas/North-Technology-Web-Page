package controllers

import (
	. "NorthTechWebPage/Log"
	. "NorthTechWebPage/Site/helpers"
	. "NorthTechWebPage/Site/models"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type AboutController struct {}

func (about AboutController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	view, err := template.ParseFiles(Include("about/")...)
	if err != nil {
		LogJson("Site","Error","About","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data:=make(map[string]interface{})
	data["about"] = About{}.GetAll()[0]
	data["menus"] = Menu{}.GetAll()
	data["features"] = Feature{}.GetAll()[0:4]

	view.ExecuteTemplate(w,"about",data)
}
