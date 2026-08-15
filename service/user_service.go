package service

import (
	"github.com/Prasaddhulgande/Crud_With_Gorm/models"
	"github.com/Prasaddhulgande/Crud_With_Gorm/config"
)
// Services doing only DB related tasks like... connect to DB, search users from DB, talking to DB

func GetAllUsers() ([]models.User, error) {
      var users []models.User                      // []models.User because of getting array of users, all users
	  result:= config.DB.Find(&users)
	  return users, result.Error
}
func GetUserByID(id string) (models.User, error) {
	var user models.User
	result:= config.DB.First(&user, id)                     //here get only one user id based
	return user, result.Error
}
func CreateUser(user *models.User) error {
	result:= config.DB.Create(user)
	return result.Error
}
func UpdateUser(user *models.User) error {
	
	result:= config.DB.Save(user)            //Save updated user data
	return result.Error
}
func DeleteUser(id string) (models.User, error) {
    var user models.User
	result := config.DB.First(&user, id)        //Search & take First result with the help of id

	if result.Error!= nil{
		return user, result.Error
	}
	del:= config.DB.Delete(&user)               //Delete user
	return user, del.Error
}