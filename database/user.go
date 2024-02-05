package database

import (
	"github.com/Mhoanganh207/BankBE/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User, db *gorm.DB) error {
	return db.Model(&models.User{}).Create(user).Error
}

func GetUser(username string, db *gorm.DB) (models.User, error) {
	result := models.User{}
	err := db.Model(&models.User{}).Where("username = ?", username).First(&result).Error
	return result, err
}
