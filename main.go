package main

import (
	"github.com/gin-gonic/gin"
	"gin-gorm-rest-docker/routes"
	"gin-gorm-rest-docker/config"
)
func main() {
    
	router := gin.Default()

	config.ConnectDB()

	routes.UserRoute(router)

	router.Run(":8000")
}