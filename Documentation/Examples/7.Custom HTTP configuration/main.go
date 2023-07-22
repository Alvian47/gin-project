package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"M": "hello",
		})
	})
	
	s := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
