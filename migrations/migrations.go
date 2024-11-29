package main

import (
  "goblogart/inits"
  "goblogart/models"
)

func init() {
  inits.LoadEnv()
  inits.DBInit()
}

func main() {
  err := inits.DB.AutoMigrate(&models.Post{})
  if err != nil {
    panic("failed to create table: " + err.Error())
  }
}
