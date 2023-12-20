package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name         string           `json:"name" form:"name"`
	Description  string           `json:"desc" form:"desc"`
	Price        int              `json:"price" form:"price"`
	Participants []Order          `json:"participants" form:"participants" gorm:"many2many:joined"`
	Submission   []SubmissionTask `json:"submission" form:"submission" gorm:"many2many:task"`
}

type EventRegister struct {
	Name         string           `json:"name" form:"name" validate:"required" `
	Description  string           `json:"desc" form:"desc" validate:"required"`
	Price        int              `json:"price" form:"price" validate:"required"`
}

type DelegasiParticipant struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Gender       string `json:"gender" form:"gender"`
	DelegasiID   uint   `json:"delegasi_id" form:"delegasi_id"`
	DelegasiType string `json:"delegasi_type" form:"delegasi_type"`
}
