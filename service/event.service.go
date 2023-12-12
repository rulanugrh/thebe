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

type eventService struct {
	repository portRepo.EventInterface
	validate   *validator.Validate
}

func NewEventServices(repository portRepo.EventInterface) portService.EventInterface {
	return &eventService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (event *eventService) Create(req domain.Event) (*web.ResponseEvent, error) {
	errValidate := middleware.ValidateStruct(event.validate, req)
	if errValidate != nil {
		log.Printf("Struct is not valid: %s", errValidate.Error())
		return nil, errValidate
	}

	data, err := event.repository.Create(req)
	if err != nil {
		log.Printf("Cant create event, because: %s", err.Error())
		errors := fmt.Sprintf("cant create, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	response := web.ResponseEvent{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
	}

	return &response, nil
}

func (event *eventService) FindByID(id uint) (interface{}, error) {
	data, err := event.repository.FindByID(id)
	if err != nil {
		log.Printf("Cant find this id, because: %s", err.Error())
		errors := fmt.Sprintf("cant find with this id, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	var listParticipant []web.ListParticipant
	var listDelegasi []web.ListDelegasi
	for _, res := range data.Participants {
		participant := web.ListParticipant{
			FName: res.UserDetail.FName,
			LName: res.UserDetail.LName,
		}

		listParticipant = append(listParticipant, participant)
		for _, resDelegasi := range res.Delegasi {
			delegasi := web.ListDelegasi{
				FName:  resDelegasi.FName,
				LName:  resDelegasi.LName,
				Gender: resDelegasi.Gender,
			}

			listDelegasi = append(listDelegasi, delegasi)
		}
	}

	if id == 2 {
		response := web.ResponseEventRekarda{
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Participant: listParticipant,
			Delegasi:    listDelegasi,
		}

		return &response, nil
	} else {
		response := web.ResponseEvent{
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Participant: listParticipant,
		}

		return &response, nil
	}

}

func (event *eventService) Update(id uint, req domain.Event) (interface{}, error) {
	data, err := event.repository.Update(id, req)
	if err != nil {
		log.Printf("Cant find this id, because: %s", err.Error())
		errors := fmt.Sprintf("cant update, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	var listParticipant []web.ListParticipant
	var listDelegasi []web.ListDelegasi
	for _, res := range data.Participants {
		participant := web.ListParticipant{
			FName: res.UserDetail.FName,
			LName: res.UserDetail.LName,
		}

		listParticipant = append(listParticipant, participant)
		for _, resDelegasi := range res.Delegasi {
			delegasi := web.ListDelegasi{
				FName:  resDelegasi.FName,
				LName:  resDelegasi.LName,
				Gender: resDelegasi.Gender,
			}

			listDelegasi = append(listDelegasi, delegasi)
		}
	}

	if id == 2 {
		response := web.ResponseEventRekarda{
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Participant: listParticipant,
			Delegasi:    listDelegasi,
		}

		return &response, nil
	} else {
		response := web.ResponseEvent{
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Participant: listParticipant,
		}

		return &response, nil
	}
}

func (event *eventService) SubmissionTask(id uint) (*web.ResponseSubmission, error) {
	data, err := event.repository.SubmissionTask(id)
	if err != nil {
		log.Printf("Cannot submission task, because: %s", err.Error())
		errors := fmt.Sprintf("cant upload task, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	response := web.ResponseSubmission{
		Name:     data.Users.FName + " " + data.Users.LName,
		Event:    data.Events.Name,
		Filename: data.Files,
	}

	return &response, nil
}
