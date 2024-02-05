package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenGenerator struct {
	SecretKey string
}

var (
	ErrorInvalidToken = errors.New("TOKEN_INVALID")
)

func NewGeneratorToken(secretKey string) Generator {
	return TokenGenerator{secretKey}
}

func (g TokenGenerator) GenerateToken(username string, duration time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.NewString(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(g.SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (g TokenGenerator) ValidateToken(jwtToken string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token parsing error: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return &claims, nil
}

func (g TokenGenerator) GetSubject(subject string) string {
	return strings.Split(subject, " ")[0]
}
