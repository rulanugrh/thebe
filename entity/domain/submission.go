package domain

import "gorm.io/gorm"

type SubmissionTask struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	EventID uint   `json:"event_ids" form:"event_ids"`
	EventDetail  Event  `json:"event" form:"event" gorm:"foreignKey:EventID;belongsTo"`
	UserID  uint   `json:"user_ids" form:"user_ids"`
	UsersDetail   User   `json:"user" form:"user" gorm:"foreignKey:UserID;belongsTo"`
	Files   string `json:"file" form:"file"`
}
