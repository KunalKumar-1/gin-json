package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalkumar-1/json-gin/initializers"
)

func init() {
	initializers.LoadEnv();
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
