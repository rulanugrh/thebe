package handler

import (
	"be-project/config"
	"be-project/entity/domain"
	"be-project/entity/web"
	portHandler "be-project/http/port"
	"be-project/middleware"
	portService "be-project/service/port"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type userHandler struct {
	service portService.UserInterface
}

func NewUserHandler(service portService.UserInterface) portHandler.UserInterface {
	return &userHandler{
		service: service,
	}
}

func (user *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)
	data, err := user.service.Register(req)
	if err != nil {
		log.Printf("Cannot register to service, because: %s", err.Error())
		response := web.WebValidationError{
			Message: "You cant register",
			Errors:  err,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		response := web.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Success register account",
			Data:    data,
		}
	
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}
	
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
func (user *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLogin
	var compare domain.User
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}
	json.Unmarshal(body, &req)

	data, err := user.service.Login(req)
	if errChck := config.DB.Where("email = ?", req.Email).First(&compare).Error; errChck != nil {
		log.Printf("cannot find with this email")
	}
	
	if err != nil {
		log.Printf("Cannot login to service, because: %s", err.Error())
		response := web.WebValidationError{
			Message: "You cant login",
			Errors: err,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		token, errToken := middleware.GenerateToken(req)
		if errToken != nil {
			response := web.ResponseFailure{
				Code:    http.StatusUnauthorized,
				Message: "You cant login without jwt",
			}
			result, errMarshalling := json.Marshal(response)
			if errMarshalling != nil {
				log.Printf("Cannot marshall response")
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
		}


		response := web.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Success login account",
			Data:    data,
		}

		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "Set-Token",
			Value:   token,
			Expires: time.Now().Add(1 * time.Hour),
			HttpOnly: true,
			Secure: false,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func (user *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)
	data, err := user.service.Update(req.Email, req)
	if err != nil {
		log.Printf("Cannot update to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "You cant update",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		response := web.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Success update account",
			Data:    data,
		}
	
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}
	
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func (user *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	err := user.service.Delete(uint(id))
	if err != nil {
		log.Printf("Cannot delete account to service, because: %s", err.Error())
		response := web.WebValidationError{
			Message: "You cant delete",
			Errors:  err,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		response := web.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Success delete account",
			Data:    "Success delete",
		}
	
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}
	
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
