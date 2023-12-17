package portService

import (
	"be-project/entity/domain"
)

type OrderInterface interface {
	Create(req domain.OrderRegister) (interface{}, error)
	Update(uuid string, req domain.Order) (interface{}, error)
	FindByUUID(uuid string) (interface{}, error)
	FindByUserID(userID uint) (interface{}, error)
	FindByUserIDDetail(uuid string, userID uint) (interface{}, error)
}