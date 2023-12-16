package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	"be-project/middleware"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository portRepo.UserRepository
	validate   *validator.Validate
	r *http.Request
}

func NewUserService(repo portRepo.UserRepository) portService.UserInterface {
	return &userService{
		repository: repo,
		validate:   validator.New(),
		r: &http.Request{},
	}
}

func (user *userService) Register(req domain.UserRegister) (*web.ResponseUser, error) {
	errValidate := middleware.ValidateStruct(user.validate, req)
	if errValidate != nil {
		log.Printf("Struct is not valid: %s", errValidate.Error())
		return nil, errValidate
	}

	hashPassword := middleware.HashPassword(req.Password)
	req.Password = hashPassword
	data, err := user.repository.Register(req)
	if err != nil {
		log.Printf("Cant register to repo user: %s", err.Error())
		errors := fmt.Sprintf("cant register, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}

	resultData := web.ResponseUser{
		ID: data.ID,
		Name:     data.Name,
		Email:     data.Email,
		Address:   data.Address,
		Telephone: data.Telephone,
		Role:      data.Role.Name,
	}

	return &resultData, nil
}
func (user *userService) Login(req domain.UserLogin) (*web.ResponseLogin, error) {
	errValidate := middleware.ValidateStruct(user.validate, req)
	if errValidate != nil {
		log.Printf("Struct is not valid: %s", errValidate.Error())
		return nil, errValidate
	}
	data, err := user.repository.FindByEmail(req)
	if err != nil {
		log.Printf("Cant find email to repo user: %s", err.Error())
		return nil, web.Error{
			Message: "email is not valid",
			Code: 401,
		}
	} 

	matchedPassword := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if matchedPassword != nil {
		return &web.ResponseLogin{}, web.Error{
			Message: "password not matched",
			Code: 401,
		}
	}
	resultData := web.ResponseLogin{
		ID: data.ID,
		Email: data.Email,
		Role: data.Role.Name,
	}

	return &resultData, nil
}

func (user *userService) Update(email string, req domain.User) (*web.ResponseUser, error) {
	data, err := user.repository.Update(email, req)
	if err != nil {
		log.Printf("Cant update user to repo user: %s", err.Error())
		errors := fmt.Sprintf("cant update, because: %s", err.Error())
		return nil, web.Error{
			Message: errors,
			Code: 400,
		}
	}
	resultData := web.ResponseUser{
		ID: data.ID,
		Name:     data.Name,
		Email:     data.Email,
		Address:   data.Address,
		Telephone: data.Telephone,
		Role:      data.Role.Name,
	}

	return &resultData, nil
}

func (user *userService) Delete(id uint) error {
	return nil
}
