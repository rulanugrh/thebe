package handler

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portHandler "be-project/http/port"
	portService "be-project/service/port"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func(user *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}
	
	json.Unmarshal(body, &req)
	data, err := user.service.Register(req)
	if err != nil {
		log.Printf("Cannot register to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code: http.StatusBadRequest,
			Message: "You cant register",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.Write(result)
		w.WriteHeader(http.StatusBadRequest)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success register account",
		Data: data,
	}

	result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

	w.Write(result)
	w.WriteHeader(http.StatusOK)

}
func(user *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}
	
	json.Unmarshal(body, &req)
	data, err := user.service.Login(req.Email)
	if err != nil {
		log.Printf("Cannot login to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code: http.StatusBadRequest,
			Message: "You cant login",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.Write(result)
		w.WriteHeader(http.StatusBadRequest)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success login account",
		Data: data,
	}

	result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

	w.Write(result)
	w.WriteHeader(http.StatusOK)
}

func(user *userHandler) Update(w http.ResponseWriter, r *http.Request) {
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
			Code: http.StatusBadRequest,
			Message: "You cant update",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.Write(result)
		w.WriteHeader(http.StatusBadRequest)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success update account",
		Data: data,
	}

	result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

	w.Write(result)
	w.WriteHeader(http.StatusOK)
}

func(user *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	err := user.service.Delete(uint(id))
	if err != nil {
		log.Printf("Cannot delete account to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code: http.StatusBadRequest,
			Message: "You cant delete",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.Write(result)
		w.WriteHeader(http.StatusBadRequest)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success delete account",
		Data: "Success delete",
	}

	result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

	w.Write(result)
	w.WriteHeader(http.StatusOK)

}