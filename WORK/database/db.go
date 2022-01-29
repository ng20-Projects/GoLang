package database

import (
	"log"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"gorm.io/gorm@v1.22.5/gorm.go"
)

var DB *gorm.DB

func GetDB() *gormDB {
	return DB
}

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname" + dbName + " sslmode=disable=disable password" + " password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate()
	DB = db
}
