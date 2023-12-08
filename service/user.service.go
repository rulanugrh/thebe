package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"log"
)

type userService struct {
	repository portRepo.UserRepository
}

func NewUserService(repo portRepo.UserRepository) portService.UserInterface {
	return &userService{
		repository: repo,
	}
}

func(user *userService) Register(req domain.User) (*web.ResponseUser, error) {
	data, err := user.repository.Register(req)
	if err != nil {
		log.Printf("Cant register to repo user: %s", err.Error())
		return nil, err
	}

	resultData := web.ResponseUser{
		FName: data.FName,
		LName: data.LName,
		Email: data.Email,
		Address: data.Address,
		Telephone: data.Telephone,
		TTL: data.TTL,
		Role: data.Role.Name,
	}

	return &resultData, nil
}
func(user *userService) Login(email string) (*web.ResponseLogin, error) {
	data, err := user.repository.FindByEmail(email)
	if err != nil {
		log.Printf("Cant find email to repo user: %s", err.Error())
		return nil, err
	}

	resultData := web.ResponseLogin {
		FName: data.FName,
		LName: data.LName,
		Email: data.Email,
	}

	return &resultData, nil
}

func(user *userService) Update(email string, req domain.User) (*web.ResponseUser, error) {
	data, err := user.repository.Update(email, req)
	if err != nil {
		log.Printf("Cant update user to repo user: %s", err.Error())
		return nil, err
	}
	resultData := web.ResponseUser{
		FName: data.FName,
		LName: data.LName,
		Email: data.Email,
		Address: data.Address,
		Telephone: data.Telephone,
		TTL: data.TTL,
		Role: data.Role.Name,
	}

	return &resultData, nil
}

func(user *userService) Delete(id uint) error {
	return nil
}