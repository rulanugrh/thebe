package portRepo

import "be-project/entity/domain"

type OrderRepository interface {
	Create(req domain.OrderRegister) (*domain.Order, error)
	Update(uuid string, req domain.Order) (*domain.Order, error)
	FindByUserID(userID uint) (*domain.Order, error)
	AppendToEvents(req domain.Payment) error
}
