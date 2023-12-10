package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UUID string `json:"uuid" form:"uuid"`
	Name string `json:"name" form:"name" `
	UserID uint `json:"user_id" form:"user_id" `
	UserDetail User `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	EventID uint `json:"event_id" form:"event_id" `
	Events Event `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
	Delegasi []DelegasiParticipant `json:"delegasi" form:"delegasi" gorm:"polymorphic:Delegasi;"`
}