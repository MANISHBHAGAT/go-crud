package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/onrush/go-crud/initializers"
	"github.com/onrush/go-crud/models"
)
        


func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		
		Title string
		Body string
  
	}

	c.Bind(&body)
	// Create a post
    post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	
	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it

	
	c.JSON(200, gin.H{
			"post": post,
		})
	}

func PostsIndex(c *gin.Context) {
		//Get the posts
	var posts []models.Post
		initializers.DB.Find(&posts)
		//Respond with them
		c.JSON(200, gin.H{
		"posts": posts,
		})
		

	}

	func PostsShow(c *gin.Context){
		
		//Get id off url
        id := c.Param("id")
		//Get the posts
		var posts models.Post
		initializers.DB.Find(&posts, id)
		//Respond with them
		c.JSON(200, gin.H{
			"posts": posts,
		})
	}

	func PostsUpdate(c *gin.Context) {
		//Get the id off the url
		id := c.Param("id")
		// Get data off req body
	var body struct {
		
		Title string
		Body string
	}

	    c.Bind(&body)
	//Find the post we're updating
	    var post models.Post
		initializers.DB.First(&post, id)
	 //Update it
        initializers.DB.Model(&post).Updates(models.Post{
			
			Title: body.Title, 
			Body: body.Body, 
		
		})

	 //Respond with it

	 c.JSON(200, gin.H{

		"post": post,
	})
	}

	func PostsDelete(c *gin.Context){
        //Get the id off the url
		id := c.Param("id")
		//Delete post
		initializers.DB.Delete(&models.Post{}, id)

		//Respond
		c.Status(200)
	}