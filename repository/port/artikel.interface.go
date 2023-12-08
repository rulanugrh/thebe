package portRepo

import "be-project/entity/domain"

type ArtikelInterface interface {
	Create(req domain.Artikel) (*domain.Artikel, error)
	FindByID(id uint) (*domain.Artikel, error)
	FindAll() ([]domain.Artikel, error)
	Delete(id uint) error
}