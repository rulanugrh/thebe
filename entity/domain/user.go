package domain

import (
	"gorm.io/gorm"
)

type User struct {
	// creating model for struct
	gorm.Model
	FName string `json:"first_name" form:"first_name" `
	LName string `json:"last_name" form:"last_name" `
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password" `
	Address string `json:"address" form:"address" `
	Telephone string `json:"telephone" form:"telephone" `
	RoleID uint `json:"role_id" form:"role_id"`
	Role Roles `json:"role" gorm:"foreignKey:RoleID;reference:ID"`
}