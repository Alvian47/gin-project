package main

import (
	"example.com/controllers"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
	
	r.Run()
}
