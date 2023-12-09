package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"log"
)

type orderService struct {
	repository portRepo.OrderRepository
}

func NewOrderService(repo portRepo.OrderRepository) portService.OrderInterface {
	return &orderService{
		repository: repo,
	}
}

func(order *orderService) Create(req domain.Order) (interface{}, error) {	
	data, err := order.repository.Create(req)
	if err != nil {
		log.Printf("Cant req order to repo, because: %s", err.Error())
		return nil, err
	}

	var listDelegasi []web.ListDelegasi
	for _ , res := range data.Delegasi {
		delegasi := web.ListDelegasi {
			FName: res.FName,
			LName: res.LName,
			Gender: res.Gender,
			
		}

		listDelegasi = append(listDelegasi, delegasi)
	}

	if req.EventID == 2 {
		resultData := web.ResponseOrderRekarda{
			UUID: data.UUID,
			FName: data.UserDetail.FName,
			LName: data.UserDetail.LName,
			Email: data.UserDetail.Email,
			Address: data.UserDetail.Address,
			Telephone: data.UserDetail.Telephone,
			EventName: data.Events.Name,
			EventPrice: data.Events.Price,
			TTL: data.UserDetail.TTL,
			Delegasi: listDelegasi,
		}

		return &resultData, nil
		
	} else {
		resultData := web.ResponseOrder{
			UUID: data.UUID,
			FName: data.UserDetail.FName,
			LName: data.UserDetail.LName,
			Email: data.UserDetail.Email,
			Address: data.UserDetail.Address,
			Telephone: data.UserDetail.Telephone,
			EventName: data.Events.Name,
			EventPrice: data.Events.Price,
			TTL: data.UserDetail.TTL,
		}
	
		return &resultData, nil
	}
}

func(order *orderService) Update(uuid string, req domain.Order) (interface{}, error) {
	data, err := order.repository.Update(uuid, req)
	if err != nil {
		log.Printf("Cant req update to repo, because: %s", err.Error())
		return nil, err
	}

	var listDelegasi []web.ListDelegasi
	for _ , res := range data.Delegasi {
		delegasi := web.ListDelegasi {
			FName: res.FName,
			LName: res.LName,
			Gender: res.Gender,
			
		}

		listDelegasi = append(listDelegasi, delegasi)
	}

	if req.EventID == 2 {
		resultData := web.ResponseOrderRekarda{
			UUID: data.UUID,
			FName: data.UserDetail.FName,
			LName: data.UserDetail.LName,
			Email: data.UserDetail.Email,
			Address: data.UserDetail.Address,
			Telephone: data.UserDetail.Telephone,
			EventName: data.Events.Name,
			EventPrice: data.Events.Price,
			TTL: data.UserDetail.TTL,
			Delegasi: listDelegasi,
		}

		return &resultData, nil
		
	} else {
		resultData := web.ResponseOrder{
			UUID: data.UUID,
			FName: data.UserDetail.FName,
			LName: data.UserDetail.LName,
			Email: data.UserDetail.Email,
			Address: data.UserDetail.Address,
			Telephone: data.UserDetail.Telephone,
			EventName: data.Events.Name,
			EventPrice: data.Events.Price,
			TTL: data.UserDetail.TTL,
		}
	
		return &resultData, nil
	}
}

func(order *orderService) FindByUserID(userid uint) (interface{}, error) {
	data, err := order.repository.FindByUserID(userid)
	if err != nil {
		log.Printf("Cant req findbyuser id to repo, because: %s", err.Error())
		return nil, err
	}

	var listDelegasi []web.ListDelegasi
	for _ , res := range data.Delegasi {
		delegasi := web.ListDelegasi {
			FName: res.FName,
			LName: res.LName,
			Gender: res.Gender,
			
		}

		listDelegasi = append(listDelegasi, delegasi)
	}

	if data.EventID == 2 {
		resultData := web.ResponseOrderRekarda{
			UUID: data.UUID,
			FName: data.UserDetail.FName,
			LName: data.UserDetail.LName,
			Email: data.UserDetail.Email,
			Address: data.UserDetail.Address,
			Telephone: data.UserDetail.Telephone,
			EventName: data.Events.Name,
			EventPrice: data.Events.Price,
			TTL: data.UserDetail.TTL,
			Delegasi: listDelegasi,
		}

		return &resultData, nil
		
	} else {
		resultData := web.ResponseOrder{
			UUID: data.UUID,
			FName: data.UserDetail.FName,
			LName: data.UserDetail.LName,
			Email: data.UserDetail.Email,
			Address: data.UserDetail.Address,
			Telephone: data.UserDetail.Telephone,
			EventName: data.Events.Name,
			EventPrice: data.Events.Price,
			TTL: data.UserDetail.TTL,
		}
	
		return &resultData, nil
	}
}