package controllers

import (
	"fmt"
	"gin_session_auth/globals"
	"log"
	"net/http"

	"gin_session_auth/helpers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get(globals.UserKey)
		if user != nil {
			ctx.HTML(http.StatusBadRequest, "login.html", gin.H{
				"content": "Please logout first",
				"user":    user,
			})
			return
		}

		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get(globals.UserKey)
		if user != nil {
			ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.UserKey, username)
		session.Save()
		if err := session.Save(); err != nil {
			ctx.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/private/dashboard")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("JALAN")
		session := sessions.Default(ctx)
		user := session.Get(globals.UserKey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}

		session.Clear()
		session.Save()
		ctx.SetCookie("session", "", -1, "/", "", false, true)
		ctx.Redirect(http.StatusMovedPermanently, "/public")
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
			"user":    user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}
