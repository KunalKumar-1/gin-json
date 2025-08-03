package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalkumar-1/json-gin/controllers"
	"github.com/kunalkumar-1/json-gin/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PATCH("/posts/:id", controllers.PostsUpdate)
	
	r.Run()
}
