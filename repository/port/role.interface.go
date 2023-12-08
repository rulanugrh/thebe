package portRepo

import "be-project/entity/domain"

type RoleRepository interface {
	Create(req domain.Roles) (*domain.Roles, error)
	FindByID(id uint) (*domain.Roles, error)
	Update(id uint, req domain.Roles) (*domain.Roles, error)
	
}