package inits

import (
  "log"
  "github.com/joho/godotenv"
)

func LoadEnv() {
  err := godotenv.Load("/Users/administrator/projects/goblogart/.env")
  if err != nil {
    log.Fatal("Error loading .env file", err)
  }
}
