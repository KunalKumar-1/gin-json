package main

import (
	"github.com/kunalkumar-1/json-gin/initializers"
	"github.com/kunalkumar-1/json-gin/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
