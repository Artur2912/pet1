package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUser struct {
	Name     *string `json:"name" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty,email" gorm:"unique"`
	Password *string `json:"password" validate:"omitempty,min=6"`
}

type ResponseUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Time  time.Time `json:"time"`
}
