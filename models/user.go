package models

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Sites   []Site
}

func (u User) Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}
