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
  inits.DB.AutoMigrate(&models.Post{})
  inits.DB.AutoMigrate(&models.User{})
  // if err != nil {
  //   panic("failed to migrate table: " + err.Error())
  // }
}
