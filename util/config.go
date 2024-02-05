package util

import (
	"os"
	"time"
)

type Config struct {
	Port            string
	SecretKey       string
	Duration        time.Duration
	RefreshDuration time.Duration
}

func LoadConfig() Config {
	duration, err := time.ParseDuration(os.Getenv("DURATION"))
	if err != nil {
		panic("Parse duration failed")
	}
	return Config{
		Port:            os.Getenv("PORT"),
		SecretKey:       os.Getenv("SECRET_KEY"),
		Duration:        duration,
		RefreshDuration: time.Hour * 24 * 7,
	}

}
