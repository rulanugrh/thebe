package portRepo

import "be-project/entity/domain"

type UserRepository interface {
	Register(req domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(email string, req domain.User) (*domain.User, error)
	Delete(id uint) error
}