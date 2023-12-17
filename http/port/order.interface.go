package portHandler

import "net/http"

type OrderInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByUUID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	FindByUserID(w http.ResponseWriter, r *http.Request)
	FindByUserIDDetail(w http.ResponseWriter, r *http.Request)
}