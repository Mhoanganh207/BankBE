package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Generator interface {
	GenerateToken(username string, duration time.Duration) (string, *jwt.RegisteredClaims, error)

	ValidateToken(token string) (*jwt.MapClaims, error)

	GetSubject(subject string) string
}
