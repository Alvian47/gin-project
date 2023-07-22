package controllers

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	result := models.DB.Order("id asc").Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := models.Book{Title: book.Title, Author: book.Author}
	result := models.DB.Create(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail to create book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func FindBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	result := models.DB.Where("id = ?", id).First(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "books not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func UpdateBook(c *gin.Context) {
	var oldBook models.Book
	var newBook models.Book
	id := c.Param("id")

	result := models.DB.Where("id = ?", id).First(&oldBook)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "books not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = models.DB.Model(&oldBook).Updates(newBook)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to update book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": oldBook,
	})
}

func DeleteBook(c *gin.Context){
	var book models.Book
	id := c.Param("id")

	result := models.DB.Where("id = ?", id).First(&book)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "books not found",
		})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"msg": "book success delete",
	})
}