package main

import (
	// "fmt"

	"first/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello",
		})
	})

	r.Run()
}
