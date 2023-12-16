package portHandler

import "net/http"

type PaymentInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	HandlingStatus(w http.ResponseWriter, r *http.Request)
	PaymentNotification(w http.ResponseWriter, r *http.Request)
}