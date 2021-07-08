package controllers

import (
	. "NorthTechWebPage/Admin/helpers"
	. "NorthTechWebPage/Admin/models"
	. "NorthTechWebPage/Log"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

type ServiceController struct {}


func (service ServiceController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	CheckPageSession(w,r)
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("service/list/")...)

	if err != nil {
		LogJson("Admin","Error","Service","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["services"] = Services{}.GetAll()
	data["alerts"] = GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (service ServiceController) AddNewService(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	CheckPageSession(w,r)

	view, err := template.ParseFiles(Include("service/add/")...)
	if err != nil {
		LogJson("Admin","Error","Service","AddNewService","Could not convert html files for go to read.",err.Error())
		return
	}

	view.ExecuteTemplate(w,"index",nil)
}

func (service ServiceController) AddService(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	name := r.FormValue("name")
	description := r.FormValue("description")

	if len(description)>180 {
		description = description[0:180]
	}

	// upload file
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("picture")
	if err != nil {
		LogJson("Admin","Error","Service","AddService","Could not get files from page.",err.Error())
	}

	f,err := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		LogJson("Admin","Error","Service","AddService","Could not opened files.",err.Error())
	}

	io.Copy(f,file)

	Services{
		Name: name,
		Description: description,
		PictureUrl: "/uploads/"+header.Filename,
	}.Add()

	SetAlert(w,r,"Kayıt Başarılı!!")
	http.Redirect(w,r,"/admin/services",http.StatusSeeOther)
}

func (service ServiceController) DeleteService(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	_service := Services{}.Get(params.ByName("id"))
	_service.Delete()
	http.Redirect(w,r,"/admin/services",http.StatusSeeOther)
}

func (service ServiceController) UpdateServiceIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	CheckPageSession(w,r)

	view, err := template.ParseFiles(Include("service/edit/")...)
	if err != nil {
		LogJson("Admin","Error","Service","UpdateServiceIndex","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{}) 
	data["service"] = Services{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w,"index",data)
}

func (service ServiceController) UpdateService(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	_service := Services{}.Get(params.ByName("id"))
	name := r.FormValue("name")
	description := r.FormValue("description")
	if len(description)>180 {
		description = description[0:180]
	}

	isSelected := r.FormValue("is_selected")
	var pictureUrl string

	if isSelected == "1" {
		// upload file
		r.ParseMultipartForm(10 << 20)
		file, header,_ := r.FormFile("picture")

		f,_ := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE, 0666)

		io.Copy(f,file)
		pictureUrl= "/uploads/"+header.Filename
		os.Remove(_service.PictureUrl)

	}else{
		pictureUrl = _service.PictureUrl
	}

	_service.Update(Services{
		Name: name,
		Description: description,
		PictureUrl: pictureUrl,
	})

	http.Redirect(w,r,"/admin/services",http.StatusSeeOther)
}
