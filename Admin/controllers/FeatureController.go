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

type FeatureController struct {}

func (feature FeatureController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	CheckPageSession(w, r)
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("feature/list/")...)
	if err != nil {
		LogJson("Admin","Error", "Feature", "Index", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["features"] = Feature{}.GetAll()
	data["alerts"] = GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (feature FeatureController) AddNewFeature(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("feature/add/")...)
	if err != nil {
		LogJson("Admin","Error", "Feature", "Add", "Could not convert html files for go to read.", err.Error())
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (feature FeatureController) AddFeature(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	featureName := r.FormValue("feature-name")
	desc := r.FormValue("feature-description")
	icon := r.FormValue("feature-icon")
	Feature{
		Name: featureName,
		Description: desc,
		IconClass: icon,
	}.Add()

	SetAlert(w, r, "Kayıt Başarılı!!")
	http.Redirect(w, r, "/admin/features", http.StatusSeeOther)
}

func (feature FeatureController) DeleteFeature(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	_feature := Feature{}.Get(params.ByName("id"))
	_feature.Delete()
	http.Redirect(w, r, "/admin/features", http.StatusSeeOther)
}

func (feature FeatureController) UpdateFeatureIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)

	view, err := template.ParseFiles(Include("feature/edit/")...)
	if err != nil {
		LogJson("Admin","Error", "Feature", "Update", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["feature"] = Feature{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (feature FeatureController) UpdateFeature(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_feature := Feature{}.Get(params.ByName("id"))
	featureName := r.FormValue("feature-name")
	desc := r.FormValue("feature-description")
	icon := r.FormValue("feature-icon")

	_feature.Update(Feature{
		Name: featureName,
		Description: desc,
		IconClass: icon,
	})
	http.Redirect(w, r, "/admin/features", http.StatusSeeOther)
}
