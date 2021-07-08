package controllers

import (
	. "NorthTechWebPage/Log"
	. "NorthTechWebPage/Site/helpers"
	. "NorthTechWebPage/Site/models"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type ContactController struct {}

func (contact ContactController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(Include("contact/")...)
	if err != nil {
		LogJson("Site","Error","Contact","Index","Could not convert html files for go to read.",err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data := make(map[string]interface{})
	data["Alert"] = GetAlert(w,r)
	data["features"] = Feature{}.GetAll()[0:4]
	data["menus"] = Menu{}.GetAll()
	view.ExecuteTemplate(w,"contact", data)
}

func (contact ContactController) ContactUs(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	Message{
		Name: name,
		Email: email,
		Phone: phone,
		Message: message,
		IsReplied: false,
	}.Add()

	//SetAlert(w,r,"Mesajınız ilgili birimimize başarı ile gönderilmiştir. En kısa sürede sizinle tekrar iletişime geçeceğiz. İyi günler dileriz.")
	SetAlert(w,r,"Your message has been successfully sent to our relevant department. We will contact you again as soon as possible. We wish you a nice day.")
	http.Redirect(w,r,"/contact",http.StatusSeeOther)
}