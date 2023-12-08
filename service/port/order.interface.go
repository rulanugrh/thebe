package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type OrderInterface interface {
	Create(req domain.Order) (*web.ResponseOrder, error)
	Update(uuid string, req domain.Order) (*web.ResponseOrder, error)
	FindByUserID(userid uint) (*web.ResponseOrder, error)
}