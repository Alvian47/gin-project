package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file - error:", err.Error())
	}
}

func serveApplication() {
	router := gin.Default()

	publicRouter := router.Group("/auth")
	publicRouter.POST("/register", controller.Register)
	publicRouter.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	router.Run()
}
