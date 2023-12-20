package domain

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	EventID uint   `json:"event_id" form:"event_id"`
	Events  Event  `json:"events" form:"events" gorm:"foreignKey:EventID;references:id"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Users   User   `json:"users" form:"users" gorm:"foreignKey:UserID;references:id"`
	File   string `json:"file" form:"file"`
}

type SubmissionTask struct {
	Name    string `json:"name" form:"name" validate:"required"`
	EventID uint   `json:"event_id" form:"event_id" validate:"required"`
	UserID  uint   `json:"user_id" form:"user_id" validate:"required"`
	File   string `json:"file" form:"file" validate:"required"`
}