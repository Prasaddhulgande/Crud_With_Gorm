package controller

import (
	
	"github.com/gin-gonic/gin"
	"github.com/Prasaddhulgande/Crud_With_Gorm/models"
	//"github.com/Prasaddhulgande/Crud_With_Gorm/config"
	//"github.com/go-playground/validator/v10"
	"github.com/Prasaddhulgande/Crud_With_Gorm/service"
)

//var validate = validator.New()
 // Now controller onaly talking to HTTP , Previosuly it talking with DB & HTTP
func GetUsers(c *gin.Context) {
	users, err:= service.GetAllUsers()
	if err!= nil {
		c.JSON(400, gin.H{"Error":"Failed to fetch error"})
		return 
	}                                     //Written code line 15 to 19 instead of 20 & 21 line. It's optional code
	// users := []models.User{}
	// config.DB.Find(&users)
	c.JSON(200, users)
}
func GetUserById(c *gin.Context) {
	//var user models.User
	id := c.Param("id")

	user, err := service.GetUserByID(id)
	if err!= nil {
		c.JSON(400, gin.H{"Error": "USER NOT FOUND"})
	}
	// if err := config.DB.First(&user, id).Error; err!=nil {
    //    c.JSON(400, gin.H{"ERROR": "USER NOT FOUND"})
	//    return 
	// }
	c.JSON(200, user)
}
func CreateUser(c *gin.Context) {
	var user models.User

	if err:= c.ShouldBindJSON(&user); err!=nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	// validations 
	// if err:= validate.Struct(user); err!=nil {
	// 	c.JSON(400, gin.H{"Error": err.Error()})
	// 	return
	// }

	if err:= service.CreateUser(&user); err!= nil {
		c.JSON(400, gin.H{"Error": "Failed to create user"})
		return 
	}
	
	c.JSON(201, user)
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
    existingUser, err :=service.GetUserByID(id)
		if err!=nil{c.JSON(400, gin.H{"Error": "User Not Found"})
		return
	}
    //Bind new data into existing user
	if err:= c.ShouldBindJSON(&existingUser); err!= nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return 
	} 
    //Validate user data
	// if err:= validate.Struct(existingUser); err!=nil{
	// 	c.JSON(400, gin.H{"Error": err.Error()})
	// 	return        
	// }
    // Save updated data
	if err:= service.UpdateUser(&existingUser); err!=nil {
		c.JSON(400, gin.H{"Error": "Failed to update"})
		return
	}
		
	c.JSON(200, existingUser)
	
}
func DeleteUser(c *gin.Context) {
	//var user models.User
	id := c.Param("id")
    
	user, err := service.DeleteUser(id) 
	if err!=nil {
		c.JSON(400, gin.H{"Error": "User Not Found"})
	}
	/*//config.DB.Where("id = ?", c.Param("id")).Delete(&user) //Optional way
     
	if err:= config.DB.First(&user, id).Error; err!=nil {
		c.JSON(400, gin.H{"error": "User Not Found For Delete"})
		return 
	}
	//Delete user
	config.DB.Delete(&user) */  

	c.JSON(200, gin.H{
	"Massage": "User deleted Successfully", 
	"Deleted User": user,
})
	
}
