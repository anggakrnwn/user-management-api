package database

import (
	"fmt"
	"log"

	"github.com/anggakrnwn/user-auth-api/config"
	"github.com/anggakrnwn/user-auth-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)

	}

	fmt.Println("database connected successfully!")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	fmt.Println("database migrated successfully!")

}
