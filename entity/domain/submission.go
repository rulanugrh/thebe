package domain

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	IDEvent uint   `json:"id_event" form:"id_event"`
	Events  Event  `json:"events" form:"events" gorm:"foreignKey:IDEvent;references:ID"`
	IDUser  uint   `json:"id_user" form:"id_user"`
	Users   User   `json:"users" form:"users" gorm:"foreignKey:IDUser;references:ID"`
	File   string `json:"file" form:"file"`
	Video string `json:"link_video" form:"link_video"`
}

type SubmissionTask struct {
	Name    string `json:"name" form:"name" validate:"required"`
	IDEvent uint   `json:"id_event" form:"id_event"`
	IDUser  uint   `json:"id_user" form:"id_user"`
	Video string `json:"link_video" form:"link_video" validate:"required"`
	File   string `json:"file" form:"file" validate:"required"`
}