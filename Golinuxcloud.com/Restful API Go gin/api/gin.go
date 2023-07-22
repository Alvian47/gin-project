package api

import (
	"fmt"
	"net/http"

	"example.com/go-gin-api/database"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome building resstful api",
	})
}

func postArticle(c *gin.Context) {
	var article database.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	res, err := database.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
}

func getArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := database.ReadArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func getArticles(c *gin.Context) {
	articles, err := database.ReadArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func putArticle(c *gin.Context) {
	var article database.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error1": err,
		})
		return
	}

	res, err := database.UpdateArticle(&article)
	if err != nil {
		fmt.Println("!!!",err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ini error karena id tidak ada",
			"ERRORS": gin.H{
				"pesan": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
}

func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "article deleted successfully",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.GET("/api/v1/articles/:id", getArticle)
	r.GET("/api/v1/articles", getArticles)
	r.POST("/api/v1/articles", postArticle)
	r.PUT("/api/v1/articles", putArticle)
	r.DELETE("/api/v1/articles/:id", deleteArticle)
	return r
}
