package portRepo

import "be-project/entity/domain"

type EventInterface interface {
	Create(req domain.EventRegister) (*domain.Event, error)
	FindByID(id uint) (*domain.Event, error)
	Update(id uint, req domain.Event) (*domain.Event, error)
	SubmissionTask(req domain.Submission) (*domain.Submission, error)
}
