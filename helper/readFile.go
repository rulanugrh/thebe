package helper

import (
	"be-project/entity/web"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func ReadFormFile(form string, path string, w http.ResponseWriter ,r *http.Request) string {
	const MAX_UPLOAD_FILE_SIZE = 2048 * 2048
	
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_FILE_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_FILE_SIZE); err != nil {
		response := web.ResponseFailure{
			Message: "File is to big",
			Code: http.StatusBadRequest,
		}

		result, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}

	file, fileHeader, errForm := r.FormFile(form)
	if errForm != nil {
		response := web.ResponseFailure{
			Message: "Cannot upload with this form",
			Code: http.StatusBadRequest,
		}

		result, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}

	errCreate := os.MkdirAll(path, os.ModePerm)
	if errCreate != nil {
		log.Printf("Cannot create folder")
	}

	fileName := fmt.Sprintf("%s/%d-%s", path, time.Now().UnixNano(), fileHeader.Filename)
	dst, err := os.Create(fileName)
	if err != nil {
		response := web.ResponseFailure{
			Message: "Cannot upload file",
			Code: http.StatusBadRequest,
		}

		result, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Printf("Cant coppy files, because: %s", err.Error())
	}

	return fileName
}