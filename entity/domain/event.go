package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name         string  `json:"name" form:"name" `
	Description  string  `json:"desc" form:"desc" `
	Price        int     `json:"price" form:"price" `
	Participants []Order `json:"participants" form:"participants" gorm:"many2many:joined"`
}

type DelegasiParticipant struct {
	gorm.Model
	FName        string `json:"first_name" form:"first_name"`
	LName        string `json:"last_name" form:"last_name"`
	Gender       string `json:"gender" form:"gender"`
	DelegasiID   uint   `json:"delegasi_id" form:"delegasi_id"`
	DelegasiType string `json:"delegasi_type" form:"delegasi_type"`
}
