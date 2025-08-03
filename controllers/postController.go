package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalkumar-1/json-gin/initializers"
	"github.com/kunalkumar-1/json-gin/models"
)

func PostsCreate(ctx *gin.Context) {
	// get data form the request body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// create the post
	post := models.Post{Title: body.Title, Body: body.Body}
	r := initializers.DB.Create(&post) // pass pointer of data to Create
	if r.Error != nil {
		ctx.Status(400)
		return
	}

	// return
	ctx.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	
	// return the posts
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}
