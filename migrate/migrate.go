package main

import (
	"github.com/onrush/go-crud/initializers"
	"github.com/onrush/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
    
}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})
}