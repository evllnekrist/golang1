package database

import (
	"fmt"
	"log"
	"photo-app/models"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbPort   = os.Getenv("PGPORT")
	dbname   = os.Getenv("PGDATABASE")
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
