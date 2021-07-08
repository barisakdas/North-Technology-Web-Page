package controllers

import (
	. "NorthTechWebPage/Admin/helpers"
	. "NorthTechWebPage/Admin/models"
	. "NorthTechWebPage/Log"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

type AboutController struct {}

func (about AboutController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	CheckPageSession(w, r)
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("about/list/")...)

	if err != nil {
		LogJson("Admin","Error", "About", "Index", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["about"] = About{}.GetAll()[0]
	data["alerts"] = GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (about AboutController) EditAbout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("about/edit/")...)
	if err != nil {
		LogJson("Admin","Error", "About", "Update", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["about"] = About{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (about AboutController) UpdateAbout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_about := About{}.Get(params.ByName("id"))
	name := r.FormValue("name")
	description := r.FormValue("description")

	_about.Update(About{
		Title: name,
		Description: description,
	})
	http.Redirect(w, r, "/admin/about-us", http.StatusSeeOther)
}

