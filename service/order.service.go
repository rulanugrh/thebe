package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	"be-project/middleware"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type orderService struct {
	repository portRepo.OrderRepository
	validate   *validator.Validate
}

func NewOrderService(repo portRepo.OrderRepository) portService.OrderInterface {
	return &orderService{
		repository: repo,
		validate:   validator.New(),
	}
}

func (order *orderService) Create(req domain.OrderRegister) (interface{}, error) {
	errValidate := middleware.ValidateStruct(order.validate, req)
	if errValidate != nil {
		log.Printf("Struct is not valid: %s", errValidate.Error())
		return nil, errValidate
	}

	data, err := order.repository.Create(req)
	if err != nil {
		log.Printf("Cant req order to repo, because: %s", err.Error())
		errors := fmt.Sprintf("cant create order, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	var listDelegasi []web.ListDelegasi
	for _, res := range data.Delegasi {
		delegasi := web.ListDelegasi{
			Name:  res.Name,
			Gender: res.Gender,
		}

		listDelegasi = append(listDelegasi, delegasi)
	}

	if req.EventID == 2 {
		resultData := web.ResponseOrderRekarda{
			UUID:       data.UUID,
			Name:      data.UserDetail.Name,
			Email:      data.UserDetail.Email,
			Address:    data.UserDetail.Address,
			Telephone:  data.UserDetail.Telephone,
			EventName:  data.Events.Name,
			EventPrice: data.Events.Price,
			Delegasi:   listDelegasi,
		}

		return &resultData, nil

	} else {
		resultData := web.ResponseOrder{
			UUID:       data.UUID,
			Name:      data.UserDetail.Name,
			Email:      data.UserDetail.Email,
			Address:    data.UserDetail.Address,
			Telephone:  data.UserDetail.Telephone,
			EventName:  data.Events.Name,
			EventPrice: data.Events.Price,
		}

		return &resultData, nil
	}
}

func (order *orderService) Update(uuid string, req domain.Order) (interface{}, error) {
	data, err := order.repository.Update(uuid, req)
	if err != nil {
		log.Printf("Cant req update to repo, because: %s", err.Error())
		errors := fmt.Sprintf("cant update, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	var listDelegasi []web.ListDelegasi
	for _, res := range data.Delegasi {
		delegasi := web.ListDelegasi{
			Name:  res.Name,
			Gender: res.Gender,
		}

		listDelegasi = append(listDelegasi, delegasi)
	}

	if req.EventID == 2 {
		resultData := web.ResponseOrderRekarda{
			UUID:       data.UUID,
			Name:      data.UserDetail.Name,
			Email:      data.UserDetail.Email,
			Address:    data.UserDetail.Address,
			Telephone:  data.UserDetail.Telephone,
			EventName:  data.Events.Name,
			EventPrice: data.Events.Price,
			Delegasi:   listDelegasi,
		}

		return &resultData, nil

	} else {
		resultData := web.ResponseOrder{
			UUID:       data.UUID,
			Name:      data.UserDetail.Name,
			Email:      data.UserDetail.Email,
			Address:    data.UserDetail.Address,
			Telephone:  data.UserDetail.Telephone,
			EventName:  data.Events.Name,
			EventPrice: data.Events.Price,
		}

		return &resultData, nil
	}
}

func (order *orderService) FindByUserID(userid string) (interface{}, error) {
	data, err := order.repository.FindByUserID(userid)
	if err != nil {
		log.Printf("Cant req findbyuser id to repo, because: %s", err.Error())
		return nil, err
	}

	var listDelegasi []web.ListDelegasi
	for _, res := range data.Delegasi {
		delegasi := web.ListDelegasi{
			Name:  res.Name,
			Gender: res.Gender,
		}

		listDelegasi = append(listDelegasi, delegasi)
	}

	if data.EventID == 2 {
		resultData := web.ResponseOrderRekarda{
			UUID:       data.UUID,
			Name:      data.UserDetail.Name,
			Email:      data.UserDetail.Email,
			Address:    data.UserDetail.Address,
			Telephone:  data.UserDetail.Telephone,
			EventName:  data.Events.Name,
			EventPrice: data.Events.Price,
			Delegasi:   listDelegasi,
		}

		return &resultData, nil

	} else {
		resultData := web.ResponseOrder{
			UUID:       data.UUID,
			Name:      data.UserDetail.Name,
			Email:      data.UserDetail.Email,
			Address:    data.UserDetail.Address,
			Telephone:  data.UserDetail.Telephone,
			EventName:  data.Events.Name,
			EventPrice: data.Events.Price,
		}

		return &resultData, nil
	}
}
