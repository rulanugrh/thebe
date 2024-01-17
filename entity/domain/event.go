package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name         string       `json:"name" form:"name"`
	Description  string       `json:"desc" form:"desc"`
	Price        int          `json:"price" form:"price"`
	Participants []Order      `json:"participants" form:"participants" gorm:"many2many:joined"`
}

type EventRegister struct {
	Name        string `json:"name" form:"name" validate:"required" `
	Description string `json:"desc" form:"desc" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
}