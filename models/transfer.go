package models

import "time"

type Transfer struct {
	Id            int `gorm:"primary_key"`
	FromAccountId int
	FromAccount   Account `gorm:"foreignkey:FromAccountId"`
	ToAccountId   int
	ToAccount     Account `gorm:"foreignkey:ToAccountId"`
	Amount        int64
	Currency      string `validate:"oneof=USD VND JPY"`
	CreatedAt     time.Time
}
