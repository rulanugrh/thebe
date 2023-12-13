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

type orderHandler struct {
	service portService.OrderInterface
}

func NewOrderHandler(service portService.OrderInterface) portHandler.OrderInterface {
	return &orderHandler{
		service: service,
	}
}

func (order *orderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Order
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)
	data, err := order.service.Create(req)
	if err != nil {
		log.Printf("Cannot create order to service, because: %s", err.Error())
		response := web.WebValidationError{
			Message: "You cant create order",
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
			Message: "Success create order",
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

func (order *orderHandler) FindByUserID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["user_id"]
	id, _ := strconv.Atoi(parameter)

	data, err := order.service.FindByUserID(uint(id))
	if err != nil {
		log.Printf("Cannot find order with this id to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "You cant find order with this user id",
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
			Message: "Success find order",
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

func (order *orderHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.Order

	getID := mux.Vars(r)
	parameter := getID["id"]

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
		data, err := order.service.Update(parameter, req)
		if err != nil {
			log.Printf("Cannot update order to service, because: %s", err.Error())
			response := web.ResponseFailure{
				Code:    http.StatusBadRequest,
				Message: "You cant update order",
				Error: err,
			}
			result, errMarshalling := json.Marshal(response)
			if errMarshalling != nil {
				log.Printf("Cannot marshall response")
			}
	
			w.Write(result)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			response := web.ResponseSuccess{
				Code:    http.StatusOK,
				Message: "Success update order",
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
