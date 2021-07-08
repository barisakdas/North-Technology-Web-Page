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

type MenuController struct {}

func (menu MenuController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("menu/list/")...)
	if err != nil {
		LogJson("Admin","Error", "Feature", "Index", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["menus"] = Menu{}.GetAll()
	data["alerts"] = GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (menu MenuController) AddNewMenu(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("menu/add/")...)
	if err != nil {
		LogJson("Admin","Error", "Menu", "AddNewMenu", "Could not convert html files for go to read.", err.Error())
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (menu MenuController) AddMenu(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := r.FormValue("menu-name")
	address :=  r.FormValue("menu-address")
	Menu{
		Name: name,
		Address: address,
	}.Add()

	SetAlert(w, r, "Kayıt Başarılı!!")
	http.Redirect(w, r, "/admin/site-menus", http.StatusSeeOther)
}

func (menu MenuController) UpdateMenuIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("menu/edit/")...)
	if err != nil {
		LogJson("Admin","Error", ":Menu", "Update", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["menu"] = Menu{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (menu MenuController) UpdateMenu(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_menu := Menu{}.Get(params.ByName("id"))
	name := r.FormValue("menu-name")
	address :=  r.FormValue("menu-address")

	_menu.Update(Menu{
		Name: name,
		Address: address,
	})
	http.Redirect(w, r, "/admin/site-menus", http.StatusSeeOther)}

func (menu MenuController) DeleteMenu(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	_menu := Menu{}.Get(params.ByName("id"))
	_menu.Delete()
	http.Redirect(w, r, "/admin/site-menus", http.StatusSeeOther)
}