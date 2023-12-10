package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"

	"gorm.io/gorm"
)

type artikelRepository struct {
	db *gorm.DB
}

func NewArtikelRepository(db *gorm.DB) portRepo.ArtikelInterface {
	return &artikelRepository{
		db: db,
	}
}

func (artikel *artikelRepository) Create(req domain.Artikel) (*domain.Artikel, error) {
	err := artikel.db.Create(&req).Error
	if err != nil {
		log.Printf("Cannot create artikel in repo: %s", err.Error())
		return nil, err
	}

	return &req, nil
}

func (artikel *artikelRepository) FindByID(id uint) (*domain.Artikel, error) {
	var models domain.Artikel
	err := artikel.db.Where("id = ?", id).Find(&models).Error
	if err != nil {
		log.Printf("Cannot find artikel by this id: %s", err.Error())
		return nil, err
	}

	return &models, nil
}

func (artikel *artikelRepository) FindAll() ([]domain.Artikel, error) {
	var artikels []domain.Artikel
	err := artikel.db.Find(artikels).Error
	if err != nil {
		log.Printf("Cannot find all artikel: %s", err.Error())
		return []domain.Artikel{}, nil
	}

	return artikels, nil
}

func (artikel *artikelRepository) Delete(id uint) error {
	var models domain.Artikel
	err := artikel.db.Where("id = ?", id).Delete(&models).Error
	if err != nil {
		log.Printf("Cannot delete artikel by this id: %s", err.Error())
		return err
	}

	return nil
}
