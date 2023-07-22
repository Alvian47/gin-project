package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("Example").(string)

		// it would print : 123345
		fmt.Println("example:", example)
	})

	r.Run()
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		// set example variable
		ctx.Set("Example", "12345")

		//before request
		ctx.Next()

		// after request
		latency := time.Since(t)
		fmt.Println(latency)

		// acces status we are sending
		status := ctx.Writer.Status()
		fmt.Println(status)
	}
}
