package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/Prasaddhulgande/Crud_With_Gorm/models"
)
   var DB *gorm.DB
func ConnectDB() {
	//port=5433 it is a docker port number
	dsn := "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"

db, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})

if err!=nil {
	panic(err)
}
db.AutoMigrate(&models.User{})
DB=db

}