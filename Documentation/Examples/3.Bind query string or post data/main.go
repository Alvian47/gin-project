package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	r.POST("/testing", startPage)
	r.Run()
}

func startPage(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)
	dateString := person.Birthday.Format("2006-01-02")
	log.Println("Tahun tanggal bulan:", dateString)
	
	c.String(200, "success")
}
