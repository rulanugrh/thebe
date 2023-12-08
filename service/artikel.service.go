package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"log"
)

type artikelService struct {
	repository portRepo.ArtikelInterface
}

func NewArtikelService(repo portRepo.ArtikelInterface) portService.ArtikelInterface {
	return &artikelService{
		repository: repo,
	}
}

func(artikel *artikelService) Create(req domain.Artikel) (*web.ResponseArtikel, error) {
	data, err := artikel.repository.Create(req)
	if err != nil {
		log.Printf("Cannot create artikel in service: %s", err.Error())
		return nil, err
	}

	response := web.ResponseArtikel{
		Title: data.Title,
		Content: data.Content,
	}

	return &response, nil
}

func(artikel *artikelService) FindByID(id uint) (*web.ResponseArtikel, error) {
	data, err := artikel.repository.FindByID(id)
	if err != nil {
		log.Printf("Cannot find artikel by this id in service: %s", err.Error())
		return nil, err
	}

	response := web.ResponseArtikel{
		Title: data.Title,
		Content: data.Content,
	}

	return &response, nil
}

func(artikel *artikelService) FindAll() ([]web.ResponseArtikel, error) {
	data, err := artikel.repository.FindAll()
	if err != nil {
		log.Printf("Cannot find artikel in service: %s", err.Error())
		return nil, err
	}

	var response []web.ResponseArtikel
	for _, artikels := range data {
		artikel := web.ResponseArtikel{
			Title: artikels.Title,
			Content: artikels.Content,
		}

		response = append(response, artikel)
	}

	return response, nil
}

func(artikel *artikelService) Delete(id uint) error {
	err := artikel.repository.Delete(id)
	if err != nil {
		log.Printf("Cannot delete artikel in service: %s", err.Error())
		return err
	}
	return nil
}