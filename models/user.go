package models

import (
	"time"
)

type User struct {
    ID        uint       `json:"id,omitempty" gorm:"primaryKey" swaggerignore:"true"`
    Name      string     `json:"name" example:"Prasad"`
    Email     string     `json:"email" example:"prasad@gmail.com"`
    CreatedAt time.Time  `json:"created_at,omitempty" swaggerignore:"true"`
    UpdatedAt time.Time  `json:"updated_at,omitempty" swaggerignore:"true"`
    DeletedAt *time.Time `json:"deleted_at,omitempty" swaggerignore:"true" gorm:"index"`
}
