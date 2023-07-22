package routes

import (
	"gin_session_auth/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(ctx *gin.RouterGroup) {
	ctx.GET("/login", controllers.LoginGetHandler())
	ctx.POST("/login", controllers.LoginPostHandler())
	ctx.GET("/", controllers.IndexGetHandler())
}

func PrivateRoutes(ctx *gin.RouterGroup) {
	ctx.GET("/dashboard", controllers.DashboardGetHandler())
	ctx.GET("/log", controllers.LogoutGetHandler())
}
