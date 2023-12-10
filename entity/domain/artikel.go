package domain

import (
	"html/template"

	"gorm.io/gorm"
)

type Artikel struct {
	gorm.Model
	Title string `json:"title" validate:"required"`
	Content template.HTML `json:"content" validate:"required"`
}