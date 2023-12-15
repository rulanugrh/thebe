package portService

import (
	"be-project/entity/domain"
)

type OrderInterface interface {
	Create(req domain.OrderRegister) (interface{}, error)
	Update(uuid string, req domain.Order) (interface{}, error)
	FindByUserID(userid uint) (interface{}, error)
}