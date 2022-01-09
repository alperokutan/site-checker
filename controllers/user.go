package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"statuscheck/models"
	"statuscheck/utils"
)

type UserController struct{}

func (c UserController) GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user []models.User
		db.Find(&user)
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c UserController) CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.LogFatal(err)
		}
		db.Create(&user)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c UserController) GetUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		params := mux.Vars(r)
		userId := params["id"]
		db.First(&user, userId)
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c UserController) UpdateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		params := mux.Vars(r)
		userId := params["id"]
		db.First(&user, userId)
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.LogFatal(err)
		}
		db.Save(&user)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c UserController) DeleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		params := mux.Vars(r)
		userId := params["id"]
		db.Delete(&user, userId)
	}
}
