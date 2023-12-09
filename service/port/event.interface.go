package portService

import (
	"be-project/entity/domain"
)

type EventInterface interface {
	Create(req domain.Event) (interface{}, error)
	FindByID(id uint) (interface{}, error)
	Update(id uint, req domain.Event) (interface{}, error)
}