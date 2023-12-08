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

type orderService struct {
	repository portRepo.OrderRepository
	validate *validator.Validate
}

func NewOrderService(repo portRepo.OrderRepository) portService.OrderInterface {
	return &orderService{
		repository: repo,
		validate: validator.New(),
	}
}

func(order *orderService) Create(req domain.Order) (*web.ResponseOrder, error) {
	errStruct := middleware.ValidateStruct(order.validate, req)
	if errStruct != nil {
		return nil, errStruct
	}
	
	data, err := order.repository.Create(req)
	if err != nil {
		log.Printf("Cant req order to repo, because: %s", err.Error())
		return nil, err
	}

	resultData := web.ResponseOrder{
		UUID: data.UUID,
		FName: data.UserDetail.FName,
		LName: data.UserDetail.LName,
		Email: data.UserDetail.Email,
		Address: data.UserDetail.Address,
		TTL: data.UserDetail.TTL,
	}

	return &resultData, nil
}

func(order *orderService) Update(uuid string, req domain.Order) (*web.ResponseOrder, error) {
	data, err := order.repository.Update(uuid, req)
	if err != nil {
		log.Printf("Cant req update to repo, because: %s", err.Error())
		return nil, err
	}

	resultData := web.ResponseOrder{
		UUID: data.UUID,
		FName: data.UserDetail.FName,
		LName: data.UserDetail.LName,
		Email: data.UserDetail.Email,
		Address: data.UserDetail.Address,
		TTL: data.UserDetail.TTL,
	}

	return &resultData, nil
}

func(order *orderService) FindByUserID(userid uint) (*web.ResponseOrder, error) {
	data, err := order.repository.FindByUserID(userid)
	if err != nil {
		log.Printf("Cant req findbyuser id to repo, because: %s", err.Error())
		return nil, err
	}

	resultData := web.ResponseOrder{
		UUID: data.UUID,
		FName: data.UserDetail.FName,
		LName: data.UserDetail.LName,
		Email: data.UserDetail.Email,
		Address: data.UserDetail.Address,
		TTL: data.UserDetail.TTL,
	}

	return &resultData, nil
}