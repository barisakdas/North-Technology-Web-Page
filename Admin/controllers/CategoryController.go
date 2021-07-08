package controllers

import (
	. "NorthTechWebPage/Admin/helpers"
	. "NorthTechWebPage/Admin/models"
	. "NorthTechWebPage/Log"
	"fmt"
	"html/template"
	"time"

	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController struct{}

func (category CategoryController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("category/list/")...)

	if err != nil {
		LogJson("Admin","Error", "Category", "Index", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["categories"] = Category{}.GetAll()
	data["alerts"] = GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (category CategoryController) AddNewCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("category/add/")...)
	if err != nil {
		LogJson("Admin","Error", "Category", "Add", "Could not convert html files for go to read.", err.Error())
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (category CategoryController) AddCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryName := r.FormValue("category-name")
	Category{
		Name: categoryName,
	}.Add()

	SetAlert(w, r, "Kayıt Başarılı!!")
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}

func (category CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cat := Category{}.Get(params.ByName("id"))
	cat.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (category CategoryController) UpdateCategoryIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("category/edit/")...)
	if err != nil {
		LogJson("Admin","Error", "Category", "Update", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["category"] = Category{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (category CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cat := Category{}.Get(params.ByName("id"))
	categoryName := r.FormValue("category-name")

	cat.Update(Category{
		Name: categoryName,
	})
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}
