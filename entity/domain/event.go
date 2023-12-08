package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name        string                `json:"name" form:"name" validate:"required"`
	Description string                `json:"desc" form:"desc" validate:"required"`
	Price 		int 				  `json:"price" form:"price" validate:"required"`
	Participant []Order   			  `json:"participant" form:"participant" gorm:"many2many:parcitipant_event"`
	Delegasi    []DelegasiParticipant `json:"delegasi" form:"delegasi" gorm:"many2many:delegasi"`
}

type DelegasiParticipant struct {
	gorm.Model
	FName string `json:"first_name" form:"first_name"`
	LName string `json:"last_name" form:"last_name"`
	Gender string `json:"gender" form:"gender"`
}