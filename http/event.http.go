package handler

import (
	"be-project/entity/domain"
	"be-project/helper"

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

type eventHandler struct {
	service portService.EventInterface
}

func NewEventHandler(service portService.EventInterface) portHandler.EventInterface {
	return &eventHandler{
		service: service,
	}
}

func (event *eventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Event

	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)

	data, err := event.service.Create(req)
	if err != nil {
		log.Printf("Cannot create event to service, because: %s", err.Error())
		response := web.WebValidationError{
			Message: "You cant create event",
			Errors:  err,
		}
		result, errMarshalling := json.Marshal(response)
		if errMarshalling != nil {
			log.Printf("Cannot marshall response")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	} else {
		response := web.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Success create event",
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

func (event *eventHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	data, err := event.service.FindByID(uint(id))
	if err != nil {
		log.Printf("Cannot find event with this id to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "You cant find event with this user id",
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
			Message: "Success find event",
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

func (event *eventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.Event

	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	body, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("Cant read body request, because: %s", errRead.Error())
	}

	json.Unmarshal(body, &req)
	data, err := event.service.Update(uint(id), req)
	if err != nil {
		log.Printf("Cannot update event to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Message: "You cant update event",
			Error:  err,
			Code: 400,
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
			Message: "Success update event",
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

func (event *eventHandler) SubmissionTask(w http.ResponseWriter, r *http.Request) {
	var req domain.SubmissionTask

	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	filesname := helper.ReadFormFile("./submission/", w, *r)
	req.Files = filesname

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)

	data, err := event.service.SubmissionTask(uint(id))
	if err != nil {
		response := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "You cant create submission task",
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
			Message: "Success submission task",
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
