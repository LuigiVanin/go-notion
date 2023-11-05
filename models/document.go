package models

type Document struct {
	Base
	Title  string `json:"title"`
	Text   string `json:"text"`
	Status string `json:"status" gorm:"default:'draft'"`
	UserId uint   `json:"user_id" `
	User   User   `gorm:"foreignKey:UserId"`
}
