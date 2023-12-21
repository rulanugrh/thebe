package domain

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	EventID uint   `json:"event_id" form:"event_id"`
	Events  Event  `json:"events" form:"events" gorm:"foreignKey:EventID;references:ID"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Users   User   `json:"users" form:"users" gorm:"foreignKey:UserID;references:ID"`
	File    string `json:"file" form:"file"`
	Video   string `json:"link_video" form:"link_video"`
}

type SubmissionTask struct {
	Name    string `json:"name" form:"name" validate:"required"`
	EventID uint   `json:"event_id" form:"event_id"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Video   string `json:"link_video" form:"link_video" validate:"required"`
	File    string `json:"file" form:"file" validate:"required"`
}
