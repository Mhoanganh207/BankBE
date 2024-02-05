package database

import (
	"github.com/Mhoanganh207/BankBE/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=titbandau dbname=bank port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connect failed")
	}
	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Country{}, &models.User{}, &models.Account{}, &models.Entry{}, &models.Transfer{}, &models.Session{})
}
