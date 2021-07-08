package main

import (
	. "NorthTechWebPage/Config"
	. "NorthTechWebPage/Site/models"
	"net/http"
)

func main()  {
	DbMigration()
	http.ListenAndServe(":8080",Routes())
}
