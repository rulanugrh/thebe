package handler

import (
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

func (role *roleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Roles
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)
	errCheck := middleware.ValidateTokenAdmin(r)
	if errCheck != nil {
		log.Printf("You cant see this, just admin, %s", errCheck.Error())
		response := web.ResponseFailure{
			Code:    http.StatusForbidden,
			Message: "You cant find role by this id",
			Error: errCheck,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(result)
	} else {
		data, err := role.service.Create(req)
		if err != nil {
			log.Printf("Cannot create role to service, because: %s", err.Error())
			response := web.WebValidationError{
				Message: "You cant create role",
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
				Message: "Success create role",
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
}

func (role *roleHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	errCheck := middleware.ValidateTokenAdmin(r)
	if errCheck != nil {
		log.Printf("You cant see this, just admin, %s", errCheck.Error())
		response := web.ResponseFailure{
			Code:    http.StatusForbidden,
			Message: "You cant find role by this id",
			Error: errCheck,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(result)
	} else {
		
		data, err := role.service.FindByID(uint(id))
		if err != nil {
			log.Printf("Cannot find role by this id to service, because: %s", err.Error())
			response := web.ResponseFailure{
				Code:    http.StatusBadRequest,
				Message: "You cant find role by this id",
				Error: err,
			}
			result, errMarshalling := json.Marshal(response)
			if errMarshalling != nil {
				log.Printf("Cannot marshall response")
			}
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
	
		} else {
			response := web.ResponseSuccess{
				Code:    http.StatusOK,
				Message: "Success find role by this id",
				Data:    data,
			}
		
			result, errMarshalling := json.Marshal(response)
			if errMarshalling != nil {
				log.Printf("Cannot marshall response")
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}


}

func (role *roleHandler) Update(w http.ResponseWriter, r *http.Request) {
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
			Code:    http.StatusBadRequest,
			Message: "You cant update roles",
			Error: err,
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
			Message: "Success update roles",
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
