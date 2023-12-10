package domain

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Users       []User `json:"user" form:"user" gorm:"many2many:participant"`
}
