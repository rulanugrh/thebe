package domain

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	Name string `json:"name" form:"name" validate:"required"`
	Descript string `json:"description" form:"description" validate:"requird"`
	Users []User `json:"user" form:"user" gorm:"many2many:participant"`
}