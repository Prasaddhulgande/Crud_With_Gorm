package models
import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	//Id int      //if we imbede gorm.Model then no need to provide again.
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`

}