package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID string `json:"order_id"`
	Orders  Order  `json:"order" gorm:"foreignKey:OrderID;reference:UUID"`
}
