package services

import (
	dto "main/common/dto"
	conn "main/config/database"
	"main/models"
)

func FetchUser(userId uint) (*models.User, error) {
	user := &models.User{}
	res := conn.Database.Where(&models.User{ID: userId}).First(&user)

	return user, res.Error
}

func UpdateUser(userId uint, data dto.UpdateUser) error {
	res := conn.Database.Model(&models.User{}).Where(&models.User{ID: userId}).Updates(data)
	return res.Error
}
