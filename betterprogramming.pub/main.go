package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"gin_session_auth/middlewares"
	globals "gin_session_auth/globals"
	routes "gin_session_auth/routes"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/public")
	routes.PublicRoutes(public)

	private := router.Group("/private")
	private.Use(middlewares.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run()
}
