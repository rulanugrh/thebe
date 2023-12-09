package repository

import (
	"be-project/entity/domain"
	portRepo "be-project/repository/port"
	"log"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) portRepo.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func(role *roleRepository) Create(req domain.Roles) (*domain.Roles, error) {
	err := role.db.Create(&req).Error
	if err != nil {
		log.Printf("Cant create role, because: %s", err.Error())
		return nil, err
	}
	return &req, nil
}

func(role *roleRepository) FindByID(id uint) (*domain.Roles, error) {
	var roles domain.Roles
	err := role.db.Preload("Users").Where("id = ?", id).Find(&roles).Error
	if err != nil {
		log.Printf("Cant find role with this id, because: %s", err.Error())	
		return nil, err
	}

	return &roles, nil
}

func(role *roleRepository) Update(id uint, req domain.Roles) (*domain.Roles, error) {
	var roleUpdate domain.Roles
	err := role.db.Model(&req).Where("id = ?").Updates(&roleUpdate).Error

	if err != nil {
		log.Printf("Cant update role with this id, because: %s", err.Error())
		return nil, err
	}

	return &roleUpdate, nil
}