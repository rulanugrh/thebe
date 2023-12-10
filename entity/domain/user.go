package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// creating model for struct
	gorm.Model
	FName string `json:"first_name" form:"first_name" validate:"required"`
	LName string `json:"last_name" form:"last_name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Address string `json:"address" form:"address" validate:"required"`
	Telephone string `json:"telephone" form:"telephone" validate:"required"`
	RoleID uint `json:"role_id" form:"role_id" validate:"required"`
	Role Roles `json:"role" gorm:"foreignKey:RoleID;reference:ID"`
	TTL time.Time `json:"ttl" validate:"required"`
}