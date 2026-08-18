// @title CRUD API with GORM
// @version 1.0
// @description This is a sample CRUD API using Gin and GORM.
// @host localhost:8000
// @BasePath /
package main

import (
    "github.com/Prasaddhulgande/Crud_With_Gorm/config"
    "github.com/Prasaddhulgande/Crud_With_Gorm/routes"
    "github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
    _ "github.com/Prasaddhulgande/Crud_With_Gorm/docs"
)


func main() {
    
	router := gin.Default()

	//Swagger route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.ConnectDB()

	routes.UserRoute(router)

	router.Run(":8000")
}