package models
import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	//Id int      //if we imbede gorm.Model then no need to provide again.
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" validate: "required, min=3"`
	Email string `json:"email" validate:""required, email`
	Password string `json:"password"`

}