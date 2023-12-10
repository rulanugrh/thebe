package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"

	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) portRepo.PaymentInterface {
	return &paymentRepository{
		db: db,
	}
}

func (payment *paymentRepository) Create(req domain.Payment) (*domain.Payment, error) {
	err := payment.db.Create(&req).Error
	if err != nil {
		log.Printf("Cannot create payments: %s", err.Error())
		return nil, err
	}

	errLookup := payment.db.Preload("Orders").Preload("Orders.Events").Preload("Orders.Users").Find(&req).Error
	if errLookup != nil {
		log.Printf("Cannot lookup orders: %s", errLookup.Error())
		return nil, errLookup
	}

	return &req, nil
}

func (payment *paymentRepository) FindByID(id uint) (*domain.Payment, error) {
	var models domain.Payment
	err := payment.db.Preload("Orders").Preload("Orders.Events").Preload("Orders.Users").Where("id = ?", id).Find(&models).Error

	if err != nil {
		log.Printf("Cannot find orders by this id: %s", err.Error())
		return nil, err
	}

	return &models, nil
}

func (payment *paymentRepository) FindAll() ([]domain.Payment, error) {
	var allPayment []domain.Payment
	err := payment.db.Preload("Orders").Preload("Orders.Events").Preload("Orders.Users").Find(&allPayment).Error

	if err != nil {
		log.Printf("Cannot find orders by this id: %s", err.Error())
		return nil, err
	}

	return allPayment, nil
}
