package portHandler

import "net/http"

type RoleInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)

}