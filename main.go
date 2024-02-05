package main

import (
	"github.com/Mhoanganh207/BankBE/api"
	"github.com/Mhoanganh207/BankBE/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	server := api.NewServer(database.InitDB())
	server.Start()

}
