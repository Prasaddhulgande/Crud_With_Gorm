package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Prasaddhulgande/Crud_With_Gorm/controller"
	"github.com/Prasaddhulgande/Crud_With_Gorm/middleware"
)

func UserRoute(router *gin.Engine){
	// Public routes
    router.POST("/signup", controller.Signup)
    router.POST("/login", controller.Login)
	
    //All Users // Protected routes
	router.GET("/users", controller.GetUsers)
	//GetUserById
	router.GET("/:id", controller.GetUserById)
	router.POST("/",controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
    
	//Admin-only routes
	admin := router.Group("/admin")
   // admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    admin.Use(middleware.AuthMiddleware()) // ✅ only check token, no role
    admin.DELETE("/users/:id", controller.DeleteUser)
    //admin.PUT("/users/:id", controller.UpdateUser)

}