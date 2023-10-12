package database

import (
	"main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func ConnectDB(url string) error {
	CreateConnection(url)
	Migrate(Database)


	return nil
}