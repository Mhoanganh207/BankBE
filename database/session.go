package database

import (
	"github.com/Mhoanganh207/BankBE/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateSession(session models.Session, db *gorm.DB) error {
	return db.Model(&models.Session{}).Create(&session).Error
}

func GetSession(id uuid.UUID, db *gorm.DB) (models.Session, error) {
	result := models.Session{}
	err := db.Model(&models.Session{}).Where("id = ?", id).First(&result).Error
	return result, err
}
