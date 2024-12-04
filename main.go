package main

import (
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

  // r.GET("/", func(c *gin.Context) {
  //   c.JSON(200, gin.H{
  //     "message": "Hello World!",
  // 	})
 	// })
 	r.POST("/", controllers.CreatePost)
 	r.GET("/", controllers.GetPosts)
 	r.GET("/:id", controllers.GetPost)
 	r.PUT("/:id", controllers.UpdatePost)
 	r.DELETE("/:id", controllers.DeletePost)

  r.Run()
}
