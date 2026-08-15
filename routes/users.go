package routes

import ("github.com/gin-gonic/gin"
"github.com/Prasaddhulgande/Crud_With_Gorm/controller" 
)

func UserRoute(router *gin.Engine){
	// BELOW is GET API WITH handler function & body 
	router.GET("hello/", func(c *gin.Context){
		c.String(200, "Hello World, Welcome to the Go")
	})  
    //All Users
	router.GET("/users", controller.GetUsers)
	//GetUserById
	router.GET("/:id", controller.GetUserById)
	router.POST("/",controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}