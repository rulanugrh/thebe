package repository

import (
	"be-project/entity/domain"
	"be-project/entity/web"
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

func (order *orderRepository) Create(req domain.OrderRegister) (*domain.Order, error) {
	var models domain.Order
	models.UUID = uuid.NewString()
	models.Name = "order-" + models.UUID
	models.EventID = req.EventID
	models.UserID = req.EventID
	models.StatusPayment = "Belum Bayar"

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
		return nil, web.Error{
			Message: "Cant append to events",
			Code: 500,
		}
	}

	return &models, nil
}

func (order *orderRepository) Update(uuid string, req domain.Order) (*domain.Order, error) {
	var updateOrder domain.Order
	updateOrder.EventID = req.Events.ID
	updateOrder.UserID = req.UserID
	updateOrder.StatusPayment = req.StatusPayment
	updateOrder.Name  = req.Name
	updateOrder.UUID = uuid
	updateOrder.ID = req.ID
	err := order.db.Model(&req).Where("uuid = ?", uuid).Updates(&updateOrder).Error

	if err != nil {
		log.Printf("Cant update order with this id: %s", err.Error())
		return nil, err
	}

	return &updateOrder, nil
}

func (order *orderRepository) AppendData(req domain.Order) error {
	errAppend := order.db.Model(&req.Events).Where("id = ?", req.EventID).Association("Participants").Append(&req)
	if errAppend != nil {
		log.Printf("Cant append data because: %s", errAppend.Error())
		return web.Error{
			Message: "Cant append to events",
			Code: 500,
		}
	}

	return nil
}

func (order *orderRepository) FindByUUID(uuid string) (*domain.Order, error) {
	var orderFind domain.Order
	err := order.db.Where("uuid = ?", uuid).Find(&orderFind).Error

	if err != nil {
		log.Printf("Cant find order with this user id: %s", err.Error())
		return nil, err
	}

	return &orderFind, nil
}

func (order *orderRepository) FindByUserID(userID uint) ([]domain.Order, error) {
	var orderFind []domain.Order
	err := order.db.Preload("UserDetail").Preload("Events").Where("user_id = ?", userID).Find(&orderFind).Error

	if err != nil {
		log.Printf("Cant find order with this user id: %s", err.Error())
		return nil, err
	}

	return orderFind, nil
}

func (order *orderRepository) FindByUserIDDetail(userID uint, uuid string) (*domain.Order, error) {
	var orderFind domain.Order
	err := order.db.Preload("UserDetail").Preload("Events").Where("user_id = ?", userID).Where("uuid = ?", uuid).Find(&orderFind).Error

	if err != nil {
		log.Printf("Cant find order with this user id: %s", err.Error())
		return nil, err
	}

	return &orderFind, nil
}