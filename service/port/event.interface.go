package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type EventInterface interface {
	Create(req domain.EventRegister) (*web.ResponseEvent, error)
	FindByID(id uint) (interface{}, error)
	Update(id uint, req domain.Event) (interface{}, error)
	SubmissionTask(req domain.Submission) (*web.ResponseSubmission, error)
}
