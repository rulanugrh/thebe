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

func(event *eventService) Create(req domain.Event) (interface{}, error) {	
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

func(event *eventService) FindByID(id uint) (interface{}, error) {
	data, err := event.repository.FindByID(id)
	if err != nil {
		log.Printf("Cant find this id, because: %s", err.Error())
		return nil, err
	}

	var listParticipant []web.ListParticipant
	var listDelegasi []web.ListDelegasi
	for _, res := range data.Participants {
		participant := web.ListParticipant{
			FName: res.UserDetail.FName,
			LName: res.UserDetail.LName,
		}

		listParticipant = append(listParticipant, participant)
		for _, resDelegasi := range res.Delegasi{
			delegasi := web.ListDelegasi {
				FName: resDelegasi.FName,
				LName: resDelegasi.LName,
				Gender: resDelegasi.Gender,
			}

			listDelegasi = append(listDelegasi, delegasi)
		}
	}

	
	if id == 2 {
		response := web.ResponseEventRekarda{
			Name: data.Name,
			Description: data.Description,
			Price: data.Price,
			FileTambahan: data.FileTambahan,
			Rundown: data.Rundown,
			Materi: data.Materi,
			Participant: listParticipant,
			Delegasi: listDelegasi,
		}
	
		return &response, nil
	} else {
		response := web.ResponseEvent{
			Name: data.Name,
			Description: data.Description,
			Price: data.Price,
			FileTambahan: data.FileTambahan,
			Rundown: data.Rundown,
			Materi: data.Materi,
			Participant: listParticipant,
			
		}
	
		return &response, nil
	}
	
}

func(event *eventService) Update(id uint, req domain.Event) (interface{}, error) {
	data, err := event.repository.Update(id, req)
	if err != nil {
		log.Printf("Cant find this id, because: %s", err.Error())
		return nil, err
	}

	var listParticipant []web.ListParticipant
	var listDelegasi []web.ListDelegasi
	for _, res := range data.Participants {
		participant := web.ListParticipant{
			FName: res.UserDetail.FName,
			LName: res.UserDetail.LName,
		}

		listParticipant = append(listParticipant, participant)
		for _, resDelegasi := range res.Delegasi{
			delegasi := web.ListDelegasi {
				FName: resDelegasi.FName,
				LName: resDelegasi.LName,
				Gender: resDelegasi.Gender,
			}

			listDelegasi = append(listDelegasi, delegasi)
		}
	}

	
	if id == 2 {
		response := web.ResponseEventRekarda{
			Name: data.Name,
			Description: data.Description,
			Price: data.Price,
			FileTambahan: data.FileTambahan,
			Rundown: data.Rundown,
			Materi: data.Materi,
			Participant: listParticipant,
			Delegasi: listDelegasi,
		}
	
		return &response, nil
	} else {
		response := web.ResponseEvent{
			Name: data.Name,
			Description: data.Description,
			Price: data.Price,
			FileTambahan: data.FileTambahan,
			Rundown: data.Rundown,
			Materi: data.Materi,
			Participant: listParticipant,
		}
	
		return &response, nil
	}
}
