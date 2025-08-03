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
	var posts []models.Post // slice to hold multiple posts
	// query the database
	initializers.DB.Find(&posts) // finds all posts

	// return the posts
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(ctx *gin.Context) {
	// get post ID from URL
	id := ctx.Param("id")

	var posts []models.Post
	initializers.DB.First(&posts, id)

	ctx.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsUpdate(ctx *gin.Context) {
	//get the post

	id := ctx.Param("id")

	// bind the req body to post
	var Body struct {
		Title string
		Body  string
	}

	ctx.Bind(&Body)

	// find the post by ID
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: Body.Title,
		Body:  Body.Body,
	})

	// return the post
	ctx.JSON(200, gin.H{
		"updatedPost": post,
	})
}

func PostsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	// resp

	ctx.JSON(200, gin.H{
		"message": id + " Post Deletesd Successfully",
	})
}
