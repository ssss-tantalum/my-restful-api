package database

import (
	"log"

	"github.com/ssss.tantalum/my-restful-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(dbUrl string) {
	Instance, dbError = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannnot connect to DB")
	}
	log.Println("Conenected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
