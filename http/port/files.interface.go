package portHandler

import "net/http"

type FileInterface interface {
	UploadSeminar(w http.ResponseWriter, r *http.Request)
	UploadRekarda(w http.ResponseWriter, r *http.Request)
	UploadPID(w http.ResponseWriter, r *http.Request)
	UploadPanitia(w http.ResponseWriter, r *http.Request)
}