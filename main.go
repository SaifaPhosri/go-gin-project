package main

import (
	"go-gin-api/config"
	"go-gin-api/models"
	"go-gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Student{})

	routes.Routes(r)

	r.Run(":8080")
}
