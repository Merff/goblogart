package main

import (
	"goblogart/middlewares"
  "goblogart/inits"
  "github.com/gin-gonic/gin"
  "goblogart/controllers"
)

func init() {
  inits.LoadEnv()
  inits.DBInit()
}

func main() {
  r := gin.Default()

 	r.POST("/", middlewares.RequireAuth, controllers.CreatePost)
 	r.GET("/", controllers.GetPosts)
 	r.GET("/:id", controllers.GetPost)
 	r.PUT("/:id", controllers.UpdatePost)
 	r.DELETE("/:id", controllers.DeletePost)

 	r.POST("/users", controllers.Signup)
 	r.POST("/login", controllers.Login)
 	r.GET("/auth", controllers.Validate)
 	r.GET("/users", controllers.GetUsers)
 	r.GET("/logout", controllers.Logout)

  r.Run()
}
