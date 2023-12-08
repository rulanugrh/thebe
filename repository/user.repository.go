package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"

	"gorm.io/gorm"
)
type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) portRepo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (user *userRepository) Register(req domain.User) (*domain.User, error) {
	err := user.db.Create(req).Error
	if err != nil {
		log.Printf("Can't create user, because: %s", err.Error())
		return nil, err
	}

	errPreload := user.db.Preload("Role").Find(&req).Error
	if errPreload != nil {
		log.Printf("Can't show the roles, because: %s", errPreload.Error())
		return nil, errPreload
	}

	errAppend := user.db.Model(&req.Role).Association("user").Append(&req)
	if errAppend != nil {
		log.Printf("Can't append user, because: %s", errAppend.Error())
	}

	return &req, nil
	
}

func (user *userRepository) FindByEmail(email string) (*domain.User, error) {
	var req domain.User
	err := user.db.Preload("Role").Where("email = ?", email).Find(req).Error
	
	if err != nil {
		log.Printf("Can't login with this email: %s", err.Error())
		return nil, err
	}

	return &req, nil
}

func (user *userRepository) Update(email string, req domain.User) (*domain.User, error) {
	var updateReq domain.User
	err := user.db.Model(req).Where("email = ?", email).Updates(&updateReq).Error
	if err != nil {
		log.Printf("Can't update user with this email: %s", err.Error())
		return nil, err
	}

	return &updateReq, nil
}

func(user *userRepository) Delete(id uint) error {
	var req domain.User
	err := user.db.Where("id = ?", id).Delete(req).Error

	if err != nil {
		log.Printf("Cant delete user with this id: %s", err.Error())
		return err
	}

	return nil
}