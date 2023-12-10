package portRepo

import "be-project/entity/domain"

type PaymentInterface interface {
	Create(req domain.Payment) (*domain.Payment, error)
	FindByID(id string) (*domain.Payment, error)
	FindAll() ([]domain.Payment, error)
	Save(req domain.Transaction) error
}
