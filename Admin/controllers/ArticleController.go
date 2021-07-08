package controllers

import (
	. "NorthTechWebPage/Admin/helpers"
	. "NorthTechWebPage/Admin/models"
	. "NorthTechWebPage/Log"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ArticleController struct {}

func (article ArticleController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	CheckPageSession(w,r)

	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory":func(categoryId int) string {
			return Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(Include("article/list/")...)
	if err != nil {
		LogJson("Admin","Error","Article","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["articles"] = Article{}.GetAll()
	data["alerts"] = GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (article ArticleController) AddNewArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	CheckPageSession(w,r)

	view, err := template.ParseFiles(Include("article/add/")...)
	if err != nil {
		LogJson("Admin","Error","article","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["categories"]=Category{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (article ArticleController) AddArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	title := r.FormValue("article-title")
	description := r.FormValue("article-description")

	slug := slug.Make(title)
	content := r.FormValue("article-content")
	categoryId,_ := strconv.Atoi(r.FormValue("article-category"))
	fileName := ""

	// upload file
	err := r.ParseMultipartForm(10 << 20)

	file, header,err := r.FormFile("article-picture")
	if err != nil {
		LogJson("Admin","Error","Article","AddArticle/FormFile","Could not get files from page.",err.Error())
		fileName = "uploads/default.png"
	}

	f,err := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		LogJson("Admin","Error","Article","AddArticle/OpenFile","Could not opened files..",err.Error())
	}
	io.Copy(f,file)

	fileName = "/uploads/"+header.Filename
	Article{
		Title: title,
		Slug: slug,
		Description: description,
		CategoryID: categoryId,
		Content: content,
		PictureUrl: fileName,
	}.Add()

	SetAlert(w,r,"Kayıt Başarılı!!")
	http.Redirect(w,r,"/admin/articles",http.StatusSeeOther)
}

func (article ArticleController) DeleteArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	_article := Article{}.Get(params.ByName("id"))
	_article.Delete()
	http.Redirect(w,r,"/admin/articles",http.StatusSeeOther)
}

func (article ArticleController) UpdateArticleIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	CheckPageSession(w,r)

	view, err := template.ParseFiles(Include("article/edit/")...)
	if err != nil {
		LogJson("Admin","Error","Dashboard","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["article"] = Article{}.Get(params.ByName("id"))
	data["categories"]= Category{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (article ArticleController) UpdateArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	_article := Article{}.Get(params.ByName("id"))
	title := r.FormValue("article-title")
	description := r.FormValue("article-description")

	slug := slug.Make(title)
	content := r.FormValue("article-content")
	categoryId,_ := strconv.Atoi(r.FormValue("article-category"))
	isSelected := r.FormValue("is_selected")

	var pictureUrl string

	if isSelected == "1" {
		// upload file
		r.ParseMultipartForm(10 << 20)
		file, header,_ := r.FormFile("article-picture")

		f,_ := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE, 0666)

		io.Copy(f,file)
		pictureUrl= "/uploads/"+header.Filename
		os.Remove(_article.PictureUrl)

	}else{
		pictureUrl = _article.PictureUrl
	}

	_article.Update(Article{
		Title: title,
		Slug: slug,
		Description: description,
		CategoryID: categoryId,
		Content: content,
		PictureUrl: pictureUrl,
	})

	http.Redirect(w,r,"/admin/articles",http.StatusSeeOther)
}