package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onrush/go-crud/controllers"
	"github.com/onrush/go-crud/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/books", controllers.BooksCreate)
	r.PUT("/books/:id", controllers.BooksUpdate)
	r.GET("/books", controllers.BooksIndex)
	r.GET("/books/:id", controllers.BooksShow)
	r.DELETE("/books/:id", controllers.BooksDelete)
	r.Run()

}