package service

import (
	"github.com/Prasaddhulgande/Crud_With_Gorm/models"
	"github.com/Prasaddhulgande/Crud_With_Gorm/repository"
)
// Service now contains only business logic

func GetAllUsers() ([]models.User, error) {
      return repository.GetAllUsers()
}
func GetUserByID(id string) (models.User, error) {
	return repository.GetUserByID(id)
}
func CreateUser(user *models.User) error {
	return repository.CreateUser(user)
}
func UpdateUser(user *models.User) error {
	
	return repository.UpdateUser(user)
}
func DeleteUser(id string) (models.User, error) {
    return repository.DeleteUser(id)
}