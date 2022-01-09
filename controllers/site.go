package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"statuscheck/models"
	"statuscheck/utils"
)

type SiteController struct{}

func (c SiteController) GetSites(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var sites []models.Site
		params := mux.Vars(r)
		userId := params["userId"]
		db.Where("user_id = ?", userId).Find(&sites)

		var msgs []string
		c := make(chan string)

		for _, site := range sites{
			go checkLink(site.Url, c)
			msgs = append(msgs, <-c)
		}

		err := json.NewEncoder(w).Encode(msgs)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link + " might be down!"
		return
	}
	fmt.Println(link, "is up!")
	c <- link + "is up!"
}

func (c SiteController) CreateSite(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		site := models.Site{}
		err := json.NewDecoder(r.Body).Decode(&site)
		if err != nil {
			utils.LogFatal(err)
		}
		db.Create(&site)
		err = json.NewEncoder(w).Encode(site)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c SiteController) GetSite(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		site := models.Site{}
		params := mux.Vars(r)
		siteId := params["id"]
		db.First(&site, siteId)
		err := json.NewEncoder(w).Encode(site)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c SiteController) UpdateSite(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		site := models.Site{}
		params := mux.Vars(r)
		siteId := params["id"]
		db.First(&site, siteId)
		err := json.NewDecoder(r.Body).Decode(&site)
		if err != nil {
			utils.LogFatal(err)
		}
		db.Save(&site)
		err = json.NewEncoder(w).Encode(site)
		if err != nil {
			utils.LogFatal(err)
		}
	}
}

func (c SiteController) DeleteSite(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		site := models.Site{}
		params := mux.Vars(r)
		siteId := params["id"]
		db.Delete(&site, siteId)
	}
}

