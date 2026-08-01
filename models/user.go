package models
import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	//Id int      //if we imbede gorm.Model then no need to provide again.
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`

}