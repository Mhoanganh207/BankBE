package util

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Port      string
	SecretKey string
	Duration  time.Duration
}

func LoadConfig() Config {
	duration, err := time.ParseDuration(os.Getenv("DURATION"))
	if err != nil {
		panic("Parse duration failed")
	}
	fmt.Println("Duration: ", duration)
	fmt.Println("Port: ", os.Getenv("PORT"))
	fmt.Println("SecretKey: ", os.Getenv("SECRET_KEY"))
	return Config{
		Port:      os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Duration:  duration,
	}

}
