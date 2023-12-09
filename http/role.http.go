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

type roleHandler struct {
	service portService.RoleInterface
}

func NewRoleHandler(service portService.RoleInterface) portHandler.RoleInterface {
	return &roleHandler{
		service: service,
	}
}

func(role *roleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Roles
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}
	
	json.Unmarshal(body, &req)
	data, err := role.service.Create(req)
	if err != nil {
		log.Printf("Cannot create role to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Message: "You cant create role",
			Code: http.StatusInternalServerError,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success create role",
		Data: data,
	}

	result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func(role *roleHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	data, err := role.service.FindByID(uint(id))
	if err != nil {
		log.Printf("Cannot find role by this id to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code: http.StatusBadRequest,
			Message: "You cant find role by this id",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success find role by this id",
		Data: data,
	}

	result, errMarshalling := json.Marshal(response)
	if errMarshalling != nil {
		log.Printf("Cannot marshall response")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func(role *roleHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.Roles
	
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)
	
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}
	
	json.Unmarshal(body, &req)
	data, err := role.service.Update(uint(id), req)
	if err != nil {
		log.Printf("Cannot update role to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code: http.StatusBadRequest,
			Message: "You cant update roles",
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "Success update roles",
		Data: data,
	}

	result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}