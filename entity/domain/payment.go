package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID string `json:"order_id" validate:"required"`
	Orders  Order  `json:"order" gorm:"foreignKey:OrderID;reference:UUID"`
}

type Transaction struct {
	Name    string `json:"name"`
	Event   string `json:"event"`
	Price   int    `json:"price"`
	SnapURL string `json:"snap_url"`
	Token   string `json:"token_snap"`
}
