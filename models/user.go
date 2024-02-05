package models

import (
	"time"
)

type User struct {
	Username  string `gorm:"primary_key"`
	Password  string
	Fullname  string
	Email     string
	CreatedAt time.Time
}
