package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name         string           `json:"name" form:"name" `
	Description  string           `json:"desc" form:"desc" `
	Price        int              `json:"price" form:"price" `
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
	Name    string `json:"name" form:"name"`
	EventID uint   `json:"event_id" form:"event_id"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Events  Event  `json:"event"`
	Users   User   `json:"user"`
	Files   string `json:"file" form:"file"`
}
