package main

import (
    "github.com/Prasaddhulgande/Crud_With_Gorm/config"
    "github.com/Prasaddhulgande/Crud_With_Gorm/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    
	router := gin.Default()

	config.ConnectDB()

	routes.UserRoute(router)

	router.Run(":8000")
}