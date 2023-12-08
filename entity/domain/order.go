package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UUID string `json:"uuid" form:"uuid"`
	Name string `json:"name" form:"name" validate:"required"`
	UserID uint `json:"user_id" form:"user_id" validate:"required"`
	UserDetail User `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	ProductID uint `json:"product_id" form:"product_id" validate:"required"`
}