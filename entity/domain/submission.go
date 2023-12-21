package domain

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	EventID uint   `json:"event_id" form:"event_id"`
	Events  Event  `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Users   User   `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	File    string `json:"file" form:"file"`
	Video   string `json:"link_video" form:"link_video"`
}
