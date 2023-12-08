package portService

import (
	"be-project/entity/domain"
	"be-project/entity/web"
)

type ArtikelInterface interface {
	Create(req domain.Artikel) (*web.ResponseArtikel, error)
	FindByID(id uint) (*web.ResponseArtikel, error)
	FindAll() ([]web.ResponseArtikel, error)
	Delete(id uint) error
}