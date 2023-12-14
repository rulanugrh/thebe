package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UUID          string                `json:"uuid" form:"uuid" validate:"required"`
	Name          string                `json:"name" form:"name"`
	UserID        uint                  `json:"user_id" form:"user_id" validate:"required"`
	UserDetail    User                  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	EventID       uint                  `json:"event_id" form:"event_id" validate:"required" `
	Events        Event                 `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
	Delegasi      []DelegasiParticipant `json:"delegasi" form:"delegasi" gorm:"polymorphic:Delegasi;"`
	StatusPayment string                `json:"status_payment" form:"status_payment"`
}
