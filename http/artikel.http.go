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

type artikelHandler struct {
	service portService.ArtikelInterface
}

func NewArtikelHandler(service portService.ArtikelInterface) portHandler.ArtikelInterface {
	return &artikelHandler{
		service: service,
	}
}

func (artikel *artikelHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Artikel
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
		data, err := artikel.service.Create(req)
		if err != nil {
			log.Printf("Cannot create artikel to service, because: %s", err.Error())
			response := web.WebValidationError{
				Message: "You cant create artikel",
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
				Message: "Success create artikel",
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

func (artikel *artikelHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	parameter := getID["id"]
	id, _ := strconv.Atoi(parameter)

	data, err := artikel.service.FindByID(uint(id))
	if err != nil {
		log.Printf("Cannot find artikel by this id to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "You cant find role by this id",
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
			Message: "Success find artikel by this id",
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

func (artikel *artikelHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := artikel.service.FindAll()
	if err != nil {
		log.Printf("Cannot find all artikel to service, because: %s", err.Error())
		response := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "You cant find all artikel",
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
			Message: "Artikel found",
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

func (artikel *artikelHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
		err := artikel.service.Delete(uint(id))
		if err != nil {
			log.Printf("Cannot delete artikel to service, because: %s", err.Error())
			response := web.ResponseFailure{
				Code:    http.StatusBadRequest,
				Message: "You cant delete",
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
				Message: "Success delete artikel",
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
}
