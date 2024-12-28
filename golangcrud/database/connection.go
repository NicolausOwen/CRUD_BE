package database

import (
	"golangcrud/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "root:@tcp(localhost:3306)/golangcrud" // username:password@tcp(localhost:3306)/namadatabase
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Payment{})
	log.Println("Database berhasil diinisialisasi dan dimigrasi!")
}
