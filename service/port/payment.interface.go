package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type PaymentInterface interface {
	Create(req domain.Payment) (*web.ResponsePayment, error)
	FindByID(id string) (*web.ResponsePayment, error)
	FindAll() ([]web.ResponsePayment, error)
}
