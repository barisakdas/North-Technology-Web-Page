package models

import (
	. "NorthTechWebPage/Log"
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Feature struct {
	gorm.Model
	Name, Description, IconClass string
}

func (feature Feature) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Site","Error", "Models/Message", "Migrate", "Database migration error", err.Error())
	}
	db.AutoMigrate(&feature)
}

func (feature Feature) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Message", "Add", "Add article error", err.Error())
	}

	db.Create(&feature)
}

func (feature Feature) Get(where ...interface{}) Feature {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Get", "Get article error", err.Error())
		return feature
	}

	db.First(&feature, where...)
	return feature
}

func (feature Feature) GetAll(where ...interface{}) []Feature {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var features []Feature
	db.Find(&features, where...)
	return features
}

func (feature Feature) Update(data Feature) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Model(&feature).Updates(data)
}

func (feature Feature) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Delete(&feature)
}

