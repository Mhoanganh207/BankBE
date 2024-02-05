package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Username     string
	User         User `gorm:"foreignKey:Username;references:Username"`
	RefreshToken string
	ClientIP     string
	UserAgent    string
	IsBlocked    bool
	ExpiresAt    time.Time
}
