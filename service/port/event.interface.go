package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type EventInterface interface {
	Create(req domain.EventRegister) (*web.ResponseEvent, error)
	FindByID(id uint) (*web.ResponseEvent, error)
	Update(id uint, req domain.Event) (*web.ResponseEvent, error)
	SubmissionTask(req domain.Submission) (*web.ResponseSubmission, error)
}
