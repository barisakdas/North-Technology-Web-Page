package models

import (
	. "NorthTechWebPage/Log"
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Name, Email, Phone, Message string
	IsReplied bool
}

func (message Message) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Site","Error", "Models/Message", "Migrate", "Database migration error", err.Error())
	}
	db.AutoMigrate(&message)
}

func (message Message) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Message", "Add", "Add article error", err.Error())
	}

	db.Create(&message)
}

func (message Message) Get(where ...interface{}) Message {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Get", "Get article error", err.Error())
		return message
	}

	db.First(&message, where...)
	return message
}

func (message Message) GetAll(where ...interface{}) []Message {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var messages []Message
	db.Find(&messages, where...)
	return messages
}

func (message Message) Update(data Message) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Model(&message).Updates(data)
}

func (message Message) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Delete(&message)
}
