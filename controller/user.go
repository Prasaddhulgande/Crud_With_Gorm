package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Prasaddhulgande/Crud_With_Gorm/models"
	"github.com/Prasaddhulgande/Crud_With_Gorm/config"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}
func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err!=nil {
       c.JSON(400, gin.H{"ERROR": "USER NOT FOUND"})
	   return 
	}
	c.JSON(200, &user)
}
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	//config.DB.Where("id = ?", c.Param("id")).Delete(&user) //Optional way

	if err:= config.DB.First(&user, id).Error; err!=nil {
		c.JSON(400, gin.H{"error": "User Not Found For Delete"})
		return 
	}
	//Delete user
	config.DB.Delete(&user)
	c.JSON(200, gin.H{
	"Massage": "User deleted Successfully", 
	"Deleted User": user,
})
	
}
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	//config.DB.Where("id = ?", c.Param("id")).First(&user) //Optional way
	if err := config.DB.First(&user, id).Error; err!=nil{
		c.JSON(400, gin.H{"ERROR": "User Not Found"})
	}
     //Bind new data
	if err:= c.BindJSON(&user); err!= nil{
		c.JSON(400, gin.H{"error": "Invalid JSON"})
	}
    // Save updated data
	config.DB.Save(&user)
		
	c.JSON(200, &user)
	
}