package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name         string           `json:"name" form:"name" validate:"required" `
	Description  string           `json:"desc" form:"desc" validate:"required"`
	Price        int              `json:"price" form:"price" validate:"required"`
	Participants []Order          `json:"participants" form:"participants" gorm:"many2many:joined"`
	Submission   []SubmissionTask `json:"submission" form:"submission" gorm:"many2many:task"`
}

type DelegasiParticipant struct {
	gorm.Model
	FName        string `json:"first_name" form:"first_name"`
	LName        string `json:"last_name" form:"last_name"`
	Gender       string `json:"gender" form:"gender"`
	DelegasiID   uint   `json:"delegasi_id" form:"delegasi_id"`
	DelegasiType string `json:"delegasi_type" form:"delegasi_type"`
}

type SubmissionTask struct {
	gorm.Model
	Name    string `json:"name" form:"name" validate:"required"`
	EventID uint   `json:"event_id" form:"event_id" validate:"required"`
	UserID  uint   `json:"user_id" form:"user_id" validate:"required"`
	Events  Event  `json:"event" gorm:"foreignKey:EventID;reference:UUID"`
	Users   User   `json:"user" gorm:"foreignKey:UserID;reference:ID"`
	Files   string `json:"file" form:"file" validate:"required"`
}
