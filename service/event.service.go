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

func (event *eventService) Create(req domain.EventRegister) (*web.ResponseEvent, error) {
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
		ID: data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
	}

	return &response, nil
}

func (event *eventService) FindByID(id uint) (*web.ResponseEvent, error) {
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
	for _, res := range data.Participants {
		participant := web.ListParticipant{
			Name: res.UserDetail.Name,
			Email: res.UserDetail.Email,
		}

		listParticipant = append(listParticipant, participant)
	}

	// var listSubmission []web.ResponseSubmission
	// for _, sub := range data.Submissions {
	// 	subs := web.ResponseSubmission {
	// 		Name: sub.Name,
	// 		Event: sub.Events.Name,
	// 		Video: sub.Video,
	// 		Filename: sub.File,
	// 	}

	// 	listSubmission = append(listSubmission, subs)
	// }

	response := web.ResponseEvent{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Participant: listParticipant,
	}

	return &response, nil
}

func (event *eventService) Update(id uint, req domain.Event) (*web.ResponseEvent, error) {
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
	for _, res := range data.Participants {
		participant := web.ListParticipant{
			Name: res.UserDetail.Name,
			Email: res.UserDetail.Email,
		}

		listParticipant = append(listParticipant, participant)
	}

	// var listSubmission []web.ResponseSubmission
	// for _, sub := range data.Submissions {
	// 	subs := web.ResponseSubmission {
	// 		Name: sub.Name,
	// 		Event: sub.Events.Name,
	// 		Video: sub.Video,
	// 		Filename: sub.File,
	// 	}

	// 	listSubmission = append(listSubmission, subs)
	// }

	response := web.ResponseEvent{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Participant: listParticipant,
	}

	return &response, nil
}

func (event *eventService) SubmissionTask(req domain.Submission) (*web.ResponseSubmission, error) {
	data, err := event.repository.SubmissionTask(req)
	if err != nil {
		log.Printf("Cannot submission task, because: %s", err.Error())
		errors := fmt.Sprintf("cant upload task, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	response := web.ResponseSubmission{
		Name:     data.Users.Name,
		Event:    data.Events.Name,
		Filename: data.File,
		Video: data.Video,
	}

	return &response, nil
}
