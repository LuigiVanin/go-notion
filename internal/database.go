package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

type User struct {
	gorm.Model
	Email string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Name string `gorm:"not null"`
}


func createConnection(url string) (*gorm.DB, error)  {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		Database = db
	}
	return db, err
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}