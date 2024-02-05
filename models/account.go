package models

import "time"

type Account struct {
	Id            int `gorm:"primary_key"`
	Owner         string
	User          User `gorm:"foreignkey:Owner"`
	Balance       int64
	Currency      string
	CountryCode   int
	Country       Country `gorm:"foreignkey:CountryCode"`
	CreatedAt     time.Time
	Entries       []Entry
	TransfersFrom []Transfer `gorm:"foreignkey:FromAccountId"`
	TransfersTo   []Transfer `gorm:"foreignkey:ToAccountId"`
}
