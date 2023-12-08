package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"log"
)

type eventService struct {
	repository portRepo.EventInterface
}

func NewEventServices(repository portRepo.EventInterface) portService.EventInterface {
	return &eventService{
		repository: repository,
	}
}

func(event *eventService) Create(req domain.Event) (*web.ResponseEvent, error) {
	data, err := event.repository.Create(req)
	if err != nil {
		log.Printf("Cant create event, because: %s", err.Error())
		return nil, err
	}

	response := web.ResponseEvent {
		Name: data.Name,
		Description: data.Description,
		Price: data.Price,
	}

	return &response, nil
}

func(event *eventService) FindByID(id uint) (*web.ResponseEvent, error) {
	data, err := event.repository.FindByID(id)
	if err != nil {
		log.Printf("Cant find this id, because: %s", err.Error())
		return nil, err
	}

	var listParticipant []web.ListParticipant
	for _, res := range data.Participant {
		participant := web.ListParticipant{
			FName: res.UserDetail.FName,
			LName: res.UserDetail.LName,
		}

		listParticipant = append(listParticipant, participant)
	}

	var listDelegasi []web.ListDelegasi
	for _, res := range data.Delegasi {
		delegasi := web.ListDelegasi {
			FName: res.FName,
			LName: res.LName,
		}
		listDelegasi = append(listDelegasi, delegasi)
	}

	response := web.ResponseEvent{
		Name: data.Name,
		Description: data.Description,
		Price: data.Price,
		Participant: listParticipant,
		Delegasi: listDelegasi,
	}

	return &response, nil
}

func(event *eventService) Update(id uint, req domain.Event) (*web.ResponseEvent, error) {
	data, err := event.repository.Update(id, req)
	if err != nil {
		log.Printf("Cant find this id, because: %s", err.Error())
		return nil, err
	}

	var listParticipant []web.ListParticipant
	for _, res := range data.Participant {
		participant := web.ListParticipant{
			FName: res.UserDetail.FName,
			LName: res.UserDetail.LName,
		}

		listParticipant = append(listParticipant, participant)
	}

	var listDelegasi []web.ListDelegasi
	for _, res := range data.Delegasi {
		delegasi := web.ListDelegasi {
			FName: res.FName,
			LName: res.LName,
		}
		listDelegasi = append(listDelegasi, delegasi)
	}

	response := web.ResponseEvent{
		Name: data.Name,
		Description: data.Description,
		Price: data.Price,
		Participant: listParticipant,
		Delegasi: listDelegasi,
	}

	return &response, nil
}
