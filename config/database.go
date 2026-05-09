package config

import (
	"fmt"
	"log"
	"os"

	"atsar/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = database
	log.Println("Database berhasil terhubung!")

	// AutoMigrate
	err = DB.AutoMigrate(&models.Role{}, &models.User{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	SeedRoles()
	SeedSuperAdmin()
}

func SeedRoles() {
	var count int64
	DB.Model(&models.Role{}).Count(&count)
	if count == 0 {
		roles := []models.Role{
			{ID: 1, Name: "Super Admin"},
			{ID: 2, Name: "Customer Service"},
		}
		DB.Create(&roles)
		log.Println("Seeder Roles berhasil dijalankan.")
	}
}

func SeedSuperAdmin() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("CRMSidaq2026"), bcrypt.DefaultCost)
		user := models.User{
			NamaLengkap:   "Administrator Utama",
			NamaPanggilan: "Admin Utama",
			Email:         "admin@sidaq.id",
			Password:      string(hashedPassword),
			RoleID:        1,
		}
		DB.Create(&user)
		log.Println("Seeder Super Admin berhasil dijalankan.")
	}
}

