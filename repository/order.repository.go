package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"
	"math/rand"
	"strconv"

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

func (order *orderRepository) Create(req domain.OrderRegister) (*domain.Order, error) {
	var models domain.Order
	models.UUID = rand.Int()
	models.Name = "order-" + strconv.Itoa(models.UUID)
	models.EventID = req.EventID
	models.UserID = req.EventID
	models.Delegasi = req.Delegasi

	err := order.db.Create(&models).Error
	if err != nil {
		log.Printf("Cant creaet order. because: %s", err.Error())
		return nil, err
	}

	errsPreload := order.db.Preload("UserDetail").Preload("Events").Find(&models).Error
	if errsPreload != nil {
		log.Printf("Cant creaet order. because: %s", errsPreload.Error())
		return nil, errsPreload
	}

	errAppend := order.db.Model(&models.Events).Association("Participants").Append(&models)
	if errAppend != nil {
		log.Printf("Cant append data because: %s", errAppend.Error())
	}

	return &models, nil
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
