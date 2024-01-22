package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/onrush/go-crud/initializers"
	"github.com/onrush/go-crud/models"
)

func BooksCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Author string
		Publisher string
	}

	c.Bind(&body)
	// Create a book
	book := models.Book{Title: body.Title, Author: body.Author, Publisher :body.Publisher}
	result := initializers.DB.Create(&book)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it as result
	c.JSON(200, gin.H{
		"book": book,
	})
}

func BooksIndex(c *gin.Context) {
	//Get the Books
	var Books []models.Book
	initializers.DB.Find(&Books)

	//Respond with them
	c.JSON(200, gin.H{
		"Books": Books,
	})
}

func BooksShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//Get the Books
	var Books models.Book
	initializers.DB.Find(&Books, id)

	//Respond with them
	c.JSON(200, gin.H{
		"Books": Books,
	})
}

func BooksUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	// Get data off req body
		var body struct {
		Title string
		Author string
		Publisher string
	}

	c.Bind(&body)
	//Find the book we're updating
	var book models.Book
	initializers.DB.First(&book, id)
	
	//Update it
	initializers.DB.Model(&book).Updates(models.Book{
	Title: body.Title, 
	Author: body.Author, 
	Publisher :body.Publisher,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"book": book,
	})
}

func BooksDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")
	
	//Delete book
	initializers.DB.Delete(&models.Book{}, id)

	//Respond
	c.Status(200)
}