package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors"`
}

func formHandler(c *gin.Context) {
	fakeForm := myForm{}
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{
		"color": fakeForm.Colors,
	})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.POST("/", formHandler)
	r.GET("/form", func(ctx *gin.Context) {
		ctx.HTML(200, "form.html", nil)
	})
	r.Run()
}
