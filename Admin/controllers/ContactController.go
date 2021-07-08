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

type ContactController struct {}

func (contact ContactController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("contact/list/")...)

	if err != nil {
		LogJson("Admin","Error", "Contact", "Index", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["contacts"] = Message{}.GetAll("is_replied = ?",false)
	data["alerts"] = GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (contact ContactController) ReplyMessage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	_contact := Message{}.Get(params.ByName("id"))

	_contact.Update(Message{
		IsReplied: true,
	})
	http.Redirect(w, r, "/admin/contacts", http.StatusSeeOther)
}

func (contact ContactController) OldContacts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	view, err := template.New("index").Funcs(template.FuncMap{
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("contact/old/")...)

	if err != nil {
		LogJson("Admin","Error", "Contact", "Index", "Could not convert html files for go to read.", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["contacts"] = Message{}.GetAll("is_replied = ?",true)
	data["alerts"] = GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (contact ContactController) UnReplyMessage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CheckPageSession(w, r)
	_contact := Message{}.Get(params.ByName("id"))

	_contact.Update(Message{
		IsReplied: false,
	})
	http.Redirect(w, r, "/admin/contacts", http.StatusSeeOther)
}