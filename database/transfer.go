package database

import (
	"github.com/Mhoanganh207/BankBE/models"
	"gorm.io/gorm"
)

func CreateTransfer(tf models.Transfer, db *gorm.DB) error {
	return db.Model(&models.Transfer{}).Create(&tf).Error
}
