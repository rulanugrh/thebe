package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name        string                `json:"name" form:"name" validate:"required"`
	Description string                `json:"desc" form:"desc" validate:"required"`
	Price 		int 				  `json:"price" form:"price" validate:"required"`
	Rundown string `json:"rundown" form:"rundown"`
	Materi string `json:"materi" form:"materi"`
	FileTambahan string `json:"file_tambahan" form:"file_tambahan"`
	Participants []Order   			  `json:"participants" form:"participants" gorm:"many2many:joined"`
}

type DelegasiParticipant struct {
	gorm.Model
	FName string `json:"first_name" form:"first_name"`
	LName string `json:"last_name" form:"last_name"`
	Gender string `json:"gender" form:"gender"`
	DelegasiID uint `json:"delegasi_id" form:"delegasi_id"`
	DelegasiType string `json:"delegasi_type" form:"delegasi_type"`
}