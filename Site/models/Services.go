package models

import (
	. "NorthTechWebPage/Log"
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Services struct {
	gorm.Model
	Name, Description, PictureUrl string
}


func (service Services) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Site","Error", "Models/Services", "Migrate", "Database migration error", err.Error())
	}
	db.AutoMigrate(&service)
}

func (service Services) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Services", "Add", "Add article error", err.Error())
	}

	db.Create(&service)
}

func (service Services) Get(where ...interface{}) Services {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Services", "Get", "Get article error", err.Error())
		return service
	}

	db.First(&service, where...)
	return service
}

func (service Services) GetAll(where ...interface{}) []Services {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Services", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var services []Services
	db.Find(&services, where...)
	return services
}

func (service Services) Update(data Services) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Services", "Update", "Update article error", err.Error())
	}

	db.Model(&service).Updates(data)
}

func (service Services) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Services", "Update", "Update article error", err.Error())
	}

	db.Delete(&service)
}

