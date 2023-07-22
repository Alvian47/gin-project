package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleById(articleID); err == nil {
			c.HTML(http.StatusOK, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(ctx *gin.Context, data gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		// respond json
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// respond xml
		ctx.XML(http.StatusOK, data["payload"])
	default:
		ctx.HTML(http.StatusOK, templateName, data)
	}
}
