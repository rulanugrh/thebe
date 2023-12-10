package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) portRepo.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (order *orderRepository) Create(req domain.Order) (*domain.Order, error) {
	req.UUID = uuid.New().String()
	req.Name = "order-" + req.UUID

	err := order.db.Create(&req).Error
	if err != nil {
		log.Printf("Cant creaet order. because: %s", err.Error())
		return nil, err
	}

	errsPreload := order.db.Preload("UserDetail").Preload("Events").Find(&req).Error
	if errsPreload != nil {
		log.Printf("Cant creaet order. because: %s", errsPreload.Error())
		return nil, errsPreload
	}

	errAppend := order.db.Model(&req.Events).Association("Participants").Append(&req)
	if errAppend != nil {
		log.Printf("Cant append data because: %s", errAppend.Error())
	}

	return &req, nil
}

func (order *orderRepository) Update(uuid string, req domain.Order) (*domain.Order, error) {
	var updateOrder domain.Order
	err := order.db.Model(&req).Where("uuid = ?", uuid).Updates(&updateOrder).Error

	if err != nil {
		log.Printf("Cant update order with this id: %s", err.Error())
		return nil, err
	}

	errsPreload := order.db.Preload("UserDetail").Preload("Events").Error
	if errsPreload != nil {
		log.Printf("Cant creaet order. because: %s", errsPreload.Error())
		return nil, errsPreload
	}

	return &updateOrder, nil
}

func (order *orderRepository) FindByUserID(userID uint) (*domain.Order, error) {
	var orderFind domain.Order
	err := order.db.Preload("UserDetail").Preload("Events").Where("user_id = ?", userID).Find(&orderFind).Error

	if err != nil {
		log.Printf("Cant find order with this user id: %s", err.Error())
		return nil, err
	}

	return &orderFind, nil
}
