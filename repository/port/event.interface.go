package portRepo

import "be-project/entity/domain"

type EventInterface interface {
	Create(req domain.Event) (*domain.Event, error)
	FindByID(id uint) (*domain.Event, error)
	Update(id uint, req domain.Event) (*domain.Event, error)
	SubmissionTask(id uint) (*domain.SubmissionTask, error)
}
