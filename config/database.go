package config

import (
	models "main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func CreateConnection(url string) (*gorm.DB, error)  {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err == nil {
		Database = db
	}
	return db, err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	return err
}

func ConnectDB(url string) {
	CreateConnection(url)
	Migrate(Database)
}