package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) portRepo.EventInterface {
	return &eventRepository{
		db: db,
	}
}

func (event *eventRepository) Create(req domain.EventRegister) (*domain.Event, error) {
	models := domain.Event{
		Name: req.Name,
		Price: req.Price,
		Description: req.Description,
	}
	err := event.db.Create(&models).Error
	if err != nil {
		log.Printf("Cannot create event, because: %s", err.Error())
		return nil, err
	}

	return &models, nil
}

func (event *eventRepository) FindByID(id uint) (*domain.Event, error) {
	var req domain.Event
	err := event.db.Preload("Participants.UserDetail").Where("id = ?", id).Find(&req).Error
	if err != nil {
		log.Printf("Cannot find event by this id, because: %s", err.Error())
	}

	return &req, nil
}

func (event *eventRepository) Update(id uint, req domain.Event) (*domain.Event, error) {
	var result domain.Event
	err := event.db.Model(&req).Where("id = ?", id).Updates(&result).Error
	if err != nil {
		log.Printf("Cant update, because: %s", err.Error())
	}

	errPreload := event.db.Preload("Participants.UserDetail").Find(&req).Error
	if errPreload != nil {
		log.Printf("Cannot preload event, because: %s", err.Error())
		return nil, errPreload
	}

	return &result, nil
}

func (event *eventRepository) SubmissionTask(id uint) (*domain.SubmissionTask, error) {
	var submission domain.SubmissionTask
	err := event.db.Create(&submission).Error
	if err != nil {
		log.Printf("Cannot create submission to db: %s", err.Error())
		return nil, err
	}

	errLoad := event.db.Preload("Events").Preload("Users").Find(&submission).Error
	if errLoad != nil {
		log.Printf("Cannot preload data, %s", errLoad.Error())
		return nil, errLoad
	}

	errAppend := event.db.Model(&submission.Events).Association("Submission").Append(&submission)
	if errAppend != nil {
		log.Printf("Cannot append data, %s", errAppend.Error())
		return nil, errLoad
	}

	return &submission, nil
}
