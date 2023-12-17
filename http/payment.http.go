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

	"github.com/gorilla/mux"
)

type paymentHandler struct {
	service portService.PaymentInterface
}

func NewPaymentHandler(service portService.PaymentInterface) portHandler.PaymentInterface {
	return &paymentHandler{
		service: service,
	}
}

func (payment *paymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Payment
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)
	data, err := payment.service.Create(req)
	if err != nil {
		log.Printf("Cannot create payments to service, because: %s", err.Error())
		response := web.WebValidationError{
			Message: "You cant create payments",
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

func (payment *paymentHandler) FindByID(w http.ResponseWriter, r *http.Request) {}
func (payment *paymentHandler) FindAll(w http.ResponseWriter, r *http.Request) {}
func (payment *paymentHandler) HandlingStatus(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]

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
		data, err := payment.service.HandlingStatus(parameter)
		if err != nil {
			log.Printf("Cannot find status  with this id, because: %s", err.Error())
			response := web.WebValidationError{
				Message: "You cant find status payments",
				Errors:  err,
			}
			result, errMarshalling := json.Marshal(response)
			if errMarshalling != nil {
				log.Printf("Cannot marshall response")
			}
	
			w.WriteHeader(500)
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
}

func (payment *paymentHandler) PaymentNotification(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)

	orderId, exist := req["order_id"].(string)
	if !exist  {
		w.WriteHeader(400)
	}

	sucess, _ := payment.service.NotificationStream(orderId)
	if sucess {
		response := web.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "handling payments",
			Data: req,
		}
	
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}
	
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

	w.WriteHeader(400)
}