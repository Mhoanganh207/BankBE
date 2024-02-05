package models

type Country struct {
	Code int `gorm:"primary_key"`
	Name string
}
