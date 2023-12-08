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

func(event *eventRepository) Create(req domain.Event) (*domain.Event, error) {
	err := event.db.Create(&req).Error
	if err != nil {
		log.Printf("Cannot create event, because: %s", err.Error())
		return nil, err
	}

	return &req, nil
}

func(event *eventRepository) FindByID(id uint) (*domain.Event, error) {
	var req domain.Event
	err := event.db.Where("id = ?", id).Find(&req).Error
	if err != nil {
		log.Printf("Cannot find event by this id, because: %s", err.Error())
	}

	return &req, nil
}

func(event *eventRepository) Update(id uint, req domain.Event) (*domain.Event, error) {
	var result domain.Event
	err := event.db.Model(&req).Where("id = ?", id).Updates(&result).Error
	if err != nil {
		log.Printf("Cant update, because: %s", err.Error())
	}

	return &result, nil
}