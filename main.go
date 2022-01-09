package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"statuscheck/controllers"
	"statuscheck/driver"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {

	db := driver.ConnectDB()
	router := mux.NewRouter()

	userController := controllers.UserController{}
	siteController := controllers.SiteController{}

	router.HandleFunc("/users", userController.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser(db)).Methods("DELETE")

	router.HandleFunc("/user/sites/{userId}", siteController.GetSites(db)).Methods("GET")
	router.HandleFunc("/sites/{id}", siteController.GetSite(db)).Methods("GET")
	router.HandleFunc("/sites", siteController.CreateSite(db)).Methods("POST")
	router.HandleFunc("/sites/{id}", siteController.UpdateSite(db)).Methods("PUT")
	router.HandleFunc("/sites/{id}", siteController.DeleteSite(db)).Methods("DELETE")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	fmt.Println("Server has been started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}