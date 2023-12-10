package portHandler

import "net/http"

type ArtikelInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
