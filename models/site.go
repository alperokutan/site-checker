package models

import (
	"gorm.io/gorm"
	"log"
)

type Site struct {
	gorm.Model
	Name   string `json:"name"`
	Url    string `json:"url"`
	UserID uint   `json:"userID"`
}

func (u Site) Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Site{})
	if err != nil {
		log.Fatal(err)
	}
}
