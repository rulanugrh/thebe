package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	"be-project/middleware"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"log"

	"github.com/go-playground/validator/v10"
)

type roleService struct {
	repository portRepo.RoleRepository
	validate *validator.Validate
}

func NewRoleService(repo portRepo.RoleRepository) portService.RoleInterface {
	return &roleService{
		repository: repo,
		validate: validator.New(),
	}
}

func(role *roleService) Create(req domain.Roles) (*web.ResponseCreateRole, error) {
	errStruct := middleware.ValidateStruct(role.validate, req)
	if errStruct != nil {
		return nil, errStruct
	}
	
	data, err := role.repository.Create(req)
	if err != nil {
		log.Printf("Cant create role to repo role: %s", err.Error())
		return nil, err
	}

	resultData := web.ResponseCreateRole{
		Name: data.Name,
		Description: data.Descript,
	}

	return &resultData, nil
}

func(role *roleService) FindByID(id uint) (*web.ResponseRole, error) {
	data, err := role.repository.FindByID(id)
	if err != nil {
		log.Printf("Cant find this role to repo roles: %s", err.Error())
		return nil, err
	}

	var users []web.ResponseUser
	for _, user := range data.Users {
		oneUser := web.ResponseUser {
			FName: user.FName,
			LName: user.LName,
			Email: user.Email,
			Address: user.Address,
			Telephone: user.Telephone,
			TTL: user.TTL,
		}

		users = append(users, oneUser)
	}

	resultData := web.ResponseRole{
		Name: data.Name,
		Description: data.Descript,
		User: users,
	}

	return &resultData, nil
}

func(role *roleService) Update(id uint, req domain.Roles) (*web.ResponseRole, error) {
	data, err := role.repository.Update(id, req)
	if err != nil {
		log.Printf("Cant update this role to repo roles: %s", err.Error())
		return nil, err
	}

	var users []web.ResponseUser
	for _, user := range data.Users {
		oneUser := web.ResponseUser {
			FName: user.FName,
			LName: user.LName,
			Email: user.Email,
			Address: user.Address,
			Telephone: user.Telephone,
			TTL: user.TTL,
		}

		users = append(users, oneUser)
	}

	resultData := web.ResponseRole{
		Name: data.Name,
		Description: data.Descript,
		User: users,
	}

	return &resultData, nil
}

