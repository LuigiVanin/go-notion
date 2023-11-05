package database

import (
	"fmt"
	"main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err == nil {
		Database = db
	}
	return db, err
}

func Migrate() error {
	err := Database.AutoMigrate(&models.User{}, &models.Document{})
	fmt.Println("PRINTING DATABASE: ", err)
	return err
}

func ConnectDB(url string) error {
	CreateConnection(url)
	Migrate()

	return nil
}
