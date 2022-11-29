package database

import (
	"fmt"
	"log"
	"photo-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5433"
	dbname   = "GLNG-KM-MyGram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("errror connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.AutoMigrate(models.User{}, models.Product{}, models.Status{}, models.Order{}, models.Items{})

}

func GetDB() *gorm.DB {
	return db
}
