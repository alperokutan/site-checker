package driver

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("SQL_URL")))
	if err != nil {
		log.Fatalf("DB conn error occured. Err: %s", err)
	}
	fmt.Println("app connected to", db.Name())
	return db
}