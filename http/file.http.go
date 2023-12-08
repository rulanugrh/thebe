package handler

import (
	"be-project/entity/web"
	portHandler "be-project/http/port"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type fileHandler struct {}

func NewFileHandler() portHandler.FileInterface{
	return &fileHandler{}
}

const MAX_FILE_SIZE = 2048 * 2048
const MAX_IMAGE_SIZE = 2048 * 2048

func(file *fileHandler) UploadSeminar(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_FILE_SIZE)
	errParse := r.ParseMultipartForm(MAX_FILE_SIZE)
	if errParse != nil {
		http.Error(w, errParse.Error(), http.StatusBadRequest)
		return
	}

	rundown, headerRundown, err := r.FormFile("rundown")
	if err != nil {
		response := web.ResponseFailure {
			Code: http.StatusInternalServerError,
			Message: "Cnnot upload rundown",
		}
		result, _ := json.Marshal(response)
		w.Write(result)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer rundown.Close()

	errCreate := os.MkdirAll("/file/seminar", os.ModePerm)
	if errCreate != nil {
		log.Printf("Cannot create folder")
	}

	rundownName := fmt.Sprintf("/file/seminar/%d-%s", time.Now().UnixNano(), filepath.Ext(headerRundown.Filename))
	dst, errUpload := os.Create(rundownName)
	if errUpload != nil {
		response := web.ResponseFailure {
			Code: http.StatusInternalServerError,
			Message: "Cnnot upload rundown",
		}
		result, _ := json.Marshal(response)
		w.Write(result)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer dst.Close()

	_, errCopy := io.Copy(dst, rundown)
	if errCopy != nil {
		log.Printf("Cannot copy file to destination: %s", errCopy.Error())
	}

	materi, headerMateri, err := r.FormFile("materi")
	if err != nil {
		response := web.ResponseFailure {
			Code: http.StatusInternalServerError,
			Message: "Cnnot upload materi",
		}
		result, _ := json.Marshal(response)
		w.Write(result)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer materi.Close()
	

	materiName := fmt.Sprintf("/file/seminar/%d-%s", time.Now().UnixNano(), filepath.Ext(headerMateri.Filename))
	dstMateri, errUploadMateri := os.Create(materiName)
	if errUploadMateri != nil {
		response := web.ResponseFailure {
			Code: http.StatusInternalServerError,
			Message: "Cnnot upload materi",
		}
		result, _ := json.Marshal(response)
		w.Write(result)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer dstMateri.Close()

	_, errCopasMateri := io.Copy(dstMateri, materi)
	if errCopasMateri != nil {
		log.Printf("Cannot copy file to destination: %s", errCopy.Error())
	}

	cv, headerCV, err := r.FormFile("cv")
	if err != nil {
		response := web.ResponseFailure {
			Code: http.StatusInternalServerError,
			Message: "Cnnot upload materi",
		}
		result, _ := json.Marshal(response)
		w.Write(result)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer materi.Close()
	

	cvName := fmt.Sprintf("/file/seminar/%d-%s", time.Now().UnixNano(), filepath.Ext(headerCV.Filename))
	dstCV, errUploadCV := os.Create(cvName)
	if errUploadCV != nil {
		response := web.ResponseFailure {
			Code: http.StatusInternalServerError,
			Message: "Cnnot upload materi",
		}
		result, _ := json.Marshal(response)
		w.Write(result)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer dstMateri.Close()

	_, errCopasCV := io.Copy(dstCV, cv)
	if errCopasCV != nil {
		log.Printf("Cannot copy file to destination: %s", errCopy.Error())
	}

	data := web.FileSeminar{
		Rundown: rundownName,
		Materi: materiName,
		CV: cvName,
	}

	response := web.ResponseSuccess {
		Code: http.StatusOK,
		Message: "success upload files",
		Data: data,
	}

	result, _ := json.Marshal(response)
	w.Write(result)
	w.WriteHeader(http.StatusOK)
}

func(file *fileHandler) UploadRekarda(w http.ResponseWriter, r *http.Request) {}

func(file *fileHandler) UploadPID(w http.ResponseWriter, r *http.Request) {}
func(file *fileHandler) UploadPanitia(w http.ResponseWriter, r *http.Request)