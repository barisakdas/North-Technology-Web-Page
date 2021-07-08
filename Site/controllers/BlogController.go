package controllers

import (
	. "NorthTechWebPage/Log"
	. "NorthTechWebPage/Site/helpers"
	. "NorthTechWebPage/Site/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type BlogController struct {}

func (blog BlogController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	view, err := template.New("blog").Funcs(template.FuncMap{
		"getCategory":func(categoryId int) string {
			return Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("blog/")...)
	if err != nil {
		LogJson("Admin","Error","Article","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data:= make(map[string]interface{})
	data["articles"] = Article{}.GetAll()
	data["menus"] = Menu{}.GetAll()
	data["features"] = Feature{}.GetAll()[0:4]
	view.ExecuteTemplate(w,"blog",data)
}

func (blog BlogController) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	view, err := template.New("detail").Funcs(template.FuncMap{
		"getCategory":func(categoryId int) string {
			return Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
		"replace": func(desc string) string {
			return strings.Replace(desc, "...", "", 1)
		},
	}).ParseFiles(Include("blog/")...)
	if err != nil {
		LogJson("Admin","Error","Article","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data:= make(map[string]interface{})
	data["article"] = Article{}.Get("slug = ?",params.ByName("Slug"))
	data["menus"] = Menu{}.GetAll()
	data["features"] = Feature{}.GetAll()[0:4]
	view.ExecuteTemplate(w,"detail",data)
}

func (blog BlogController) ArticleByCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	view, err := template.New("blog").Funcs(template.FuncMap{
		"getCategory":func(categoryId int) string {
			return Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("blog/")...)
	if err != nil {
		LogJson("Admin","Error","Article","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data:= make(map[string]interface{})
	data["articles"] = Article{}.GetAll("category_id = ?",params.ByName("CategoryID"))
	data["menus"] = Menu{}.GetAll()
	data["features"] = Feature{}.GetAll()[0:4]
	//data["alerts"] = GetAlert(w,r)
	view.ExecuteTemplate(w,"blog",data)
}