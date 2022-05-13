package database

import (
	"RESTSAMPLE/models"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disbale password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}
