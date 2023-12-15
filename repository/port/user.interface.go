package portRepo

import "be-project/entity/domain"

type UserRepository interface {
	Register(req domain.UserRegister) (*domain.User, error)
	FindByEmail(req domain.UserLogin) (*domain.User, error)
	Update(email string, req domain.User) (*domain.User, error)
	Delete(id uint) error
}