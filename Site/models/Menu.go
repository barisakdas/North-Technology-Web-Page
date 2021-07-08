package models

import (
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name, Address string
}

func (menu Menu) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Menu", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&menu)
}

func (menu Menu) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Menu", "Add", "Add article error", err.Error())
	}

	db.Create(&menu)
}

func (menu Menu) Get(where ...interface{}) Menu {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Menu", "Get", "Get article error", err.Error())
		return menu
	}

	db.First(&menu, where...)
	return menu
}

func (menu Menu) GetAll(where ...interface{}) []Menu {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Menu", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var menus []Menu
	db.Find(&menus, where...)
	return menus
}

func (menu Menu) Update(data Menu) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Menu", "Update", "Update article error", err.Error())
	}

	db.Model(&menu).Updates(data)
}

func (menu Menu) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Menu", "Update", "Update article error", err.Error())
	}

	db.Delete(&menu)
}
