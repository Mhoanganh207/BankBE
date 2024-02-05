package models

import "time"

type Entry struct {
	Id        int `gorm:"primary_key"`
	Amount    int64
	AccountId int
	Account   Account `gorm:"foreignkey:AccountId"`
	CreatedAt time.Time
}
