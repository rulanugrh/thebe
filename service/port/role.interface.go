package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type RoleInterface interface {
	Create(req domain.Roles) (*web.ResponseRole, error)
	FindByID(id uint) (*web.ResponseRole, error)
	Update(id uint, req domain.Roles) (*web.ResponseRole, error)
}