package database

import (
	"github.com/Mhoanganh207/BankBE/models"
	"gorm.io/gorm"
)

func CreateAccount(account *models.Account, db *gorm.DB) error {
	return db.Model(&models.Account{}).Create(account).Error
}

func GetAccount(owner string, db *gorm.DB) (models.Account, error) {
	result := models.Account{}
	err := db.Model(&models.Account{}).Where("owner = ?", owner).First(&result).Error
	return result, err
}

func GetAccountById(id int, db *gorm.DB) (models.Account, error) {
	result := models.Account{}
	err := db.Model(&models.Account{}).Where("id = ?", id).First(&result).Error
	return result, err
}
