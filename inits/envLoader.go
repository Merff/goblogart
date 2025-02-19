package inits

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load("/Users/administrator/projects/goblogart/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
