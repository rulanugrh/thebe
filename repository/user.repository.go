package repository

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) portRepo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (user *userRepository) Register(req domain.UserRegister) (*domain.User, error) {
	models := domain.User {
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Telephone: req.Telephone,
		Address: req.Address,
		RoleID: 2,
	}

	errFind := user.db.Where("email = ?", req.Email).Find(&models).Error
	if errFind != nil {
		return nil, web.Error{
			Message: "Cant create because email has been used",
			Code: 400,
		}
		
	} else {
		err := user.db.Create(&models).Error
		if err != nil {
			log.Printf("Can't create user, because: %s", err.Error())
			return nil, err
		}

		errPreload := user.db.Preload("Role").Find(&models).Error
		if errPreload != nil {
			log.Printf("Can't show the roles, because: %s", errPreload.Error())
			return nil, errPreload
		}

		errAppend := user.db.Model(&models.Role).Association("Users").Append(&models)
		if errAppend != nil {
			log.Printf("Can't append user, because: %s", errAppend.Error())
		}

		return &models, nil
	}


}

func (user *userRepository) FindByEmail(req domain.UserLogin) (*domain.User, error) {
	var userDomain domain.User

	err := user.db.Preload("Role").Where("email = ?", req.Email).First(&userDomain).Error
	
	if err != nil {
		log.Printf("Can't login with this email or invalid password: %s", err.Error())
		return nil, web.Error{
			Message: "Cant find with this email",
			Code: 400,
		}
	}

	req.ID = userDomain.ID
	req.Role = userDomain.Role.Name

	return &userDomain, nil
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

func (user *userRepository) Delete(id uint) error {
	var req domain.User
	err := user.db.Where("id = ?", id).Delete(req).Error

	if err != nil {
		log.Printf("Cant delete user with this id: %s", err.Error())
		return err
	}

	return nil
}
