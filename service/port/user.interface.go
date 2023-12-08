package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type UserInterface interface {
	Register(req domain.User) (*web.ResponseUser, error)
	Login(email string) (*web.ResponseLogin, error)
	Update(email string, req domain.User) (*web.ResponseUser, error)
	Delete(id uint) error

}