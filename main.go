package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalkumar-1/json-gin/controllers"
	"github.com/kunalkumar-1/json-gin/initializers"
)

func init() {
	initializers.LoadEnv();
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)

	r.Run()
}
