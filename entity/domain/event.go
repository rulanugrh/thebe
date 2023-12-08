package domain

import "gorm.io/gorm"

type EventEntity struct {
	gorm.Model
	Name        string              `json:"name" form:"name" validate:"required"`
	Description string              `json:"desc" form:"desc" validate:"required"`
	Participant []ParticipantEntity `json:"participant" form:"participant" gorm:"many2many:parcitipant_event"`
	Comments    []DelegasiParticipant     `json:"delegasi" form:"delegasi" gorm:"many2many:delegasi"`
}

type ParticipantEntity struct {
	gorm.Model
	UserID  uint        `json:"user_id" form:"user_id"`
	EventID uint        `json:"event_id" form:"event_id"`
	Users   User  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   EventEntity `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}

type DelegasiParticipant struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}