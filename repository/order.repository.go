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

func(order *orderRepository) Create(req domain.Order) (*domain.Order, error) {
	req.UUID = uuid.New().String()

	err := order.db.Create(req).Error
	if err != nil {
		log.Printf("Cant creaet order. because: %s", err.Error())
		return nil, err
	}

	return &req, nil
}

func(order *orderRepository) Update(uuid string, req domain.Order) (*domain.Order, error) {
	var updateOrder domain.Order
	err := order.db.Model(&req).Where("uuid = ?", uuid).Updates(&updateOrder).Error
	
	if err != nil {
		log.Printf("Cant update order with this id: %s", err.Error())
		return nil, err
	}

	return &updateOrder, nil
}

func(order *orderRepository) FindByUserID(userID uint) (*domain.Order, error) {
	var orderFind domain.Order
	err := order.db.Preload("UserDetail").Where("user_id = ?", userID).Find(&orderFind).Error

	if err != nil {
		log.Printf("Cant find order with this user id: %s", err.Error())
		return nil, err
	}

	return &orderFind, nil
}