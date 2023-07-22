package main

import (
	"example.com/go-gin-api/api"
	"example.com/go-gin-api/database"
)

func init(){
	database.NewPostgreSQLClient()
}

func main() {
	r := api.SetupRouter()
	r.Run(":5000")
}