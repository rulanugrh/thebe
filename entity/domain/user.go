package domain

import (
	"gorm.io/gorm"
)

type User struct {
	// creating model for struct
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Address   string `json:"address" form:"address"`
	Telephone string `json:"telephone" form:"telephone"`
	RoleID    uint   `json:"role_id" form:"role_id"`
	Role      Roles  `json:"role" gorm:"foreignKey:RoleID;reference:ID"`
}


type UserLogin struct {
	ID uint `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required"`
	Role string `json:"role" form:"role"`
}

type UserRegister struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required,min=8"`
	Address   string `json:"address" form:"address" validate:"required"`
	Telephone string `json:"telephone" form:"telephone" validate:"required"`
}