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

	r.POST("/books", controllers.PostsCreate)
	r.PUT("/books/:id", controllers.PostsUpdate)
	r.GET("/books", controllers.PostsIndex)
	r.GET("/books/:id", controllers.PostsShow)
	r.DELETE("/books/:id", controllers.PostsDelete)
	r.Run()

}