package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type PaymentInterface interface {
	Create(req domain.Payment) (*web.ResponsePayment, error)
	FindByID(id string) (*web.ResponseForPayment, error)
	FindAll() ([]web.ResponseForPayment, error)
	HandlingStatus( id string ) (*web.StatusPayment, error)
	NotificationStream(orderID string) (bool, error)
}
