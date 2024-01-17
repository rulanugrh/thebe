package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type OrderInterface interface {
	Create(req domain.OrderRegister) (*web.ResponseOrder, error)
	Update(uuid string, req domain.Order) (*web.ResponseOrder, error)
	FindByUUID(uuid string) (*web.ResponseOrder, error)
	FindByUserID(userID uint) (*web.ResponseOrder, error)
	FindByUserIDDetail(uuid string, userID uint) (*web.ResponseOrder, error)
}