package models

import (
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type About struct {
	gorm.Model
	Title,Description string

}

func (about About) About()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/About", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&about)
}

func (about About) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/About", "Add", "Add article error", err.Error())
	}

	db.Create(&about)
}

func (about About) Get(where ...interface{}) About {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/About", "Get", "Get article error", err.Error())
		return about
	}

	db.First(&about, where...)
	return about
}

func (about About) GetAll(where ...interface{}) []About {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/About", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var abouts []About
	db.Find(&abouts, where...)
	return abouts
}

func (about About) Update(data About) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/About", "Update", "Update article error", err.Error())
	}

	db.Model(&about).Updates(data)
}

func (about About) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/About", "Update", "Update article error", err.Error())
	}

	db.Delete(&about)
}
