package controllers

import (
	. "NorthTechWebPage/Admin/helpers"
	. "NorthTechWebPage/Admin/models"
	. "NorthTechWebPage/Log"
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type LoginController struct {}

func (login LoginController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	view,err := template.ParseFiles(Include("login/")...)
	if err != nil {
		LogJson("Admin","Error","Login","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["alert"] = GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (login LoginController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	userName := r.FormValue("username")
	password := fmt.Sprintf("%x",sha256.Sum256([]byte(r.FormValue("password"))))

	user := User{}.Get("password = ? AND user_name = ?",password,userName)

	if user.UserName == userName && user.Password == password {
		SetUserSession(w,r,	userName,password)
		SetAlert(w,r,"Hoşgeldiniz Sn. "+user.FirstName+" "+user.LastName)
		http.Redirect(w,r,"/admin/articles",http.StatusSeeOther)
	}else{
		SetAlert(w,r,"Kullanıcı adınız veya parolanız hatalı!")
		http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
	}
}

func (login LoginController) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	RemoveUserSession(w,r)
	SetAlert(w,r,"Çıkış işlemi Başarılı!")
	http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
}