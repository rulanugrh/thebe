package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func ReadFormFile(Form string, Path string, w http.ResponseWriter ,r http.Request) string {
	const MAX_UPLOAD_FILE_SIZE = 16 * 1024 * 1024
	
	if err := r.ParseMultipartForm(MAX_UPLOAD_FILE_SIZE); err != nil {
		log.Printf("Cannot upload because file to big, %s", err.Error())
	}

	buf := make([]byte, 1024)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	errCreate := os.MkdirAll(Path, os.ModePerm)
	if errCreate != nil {
		log.Printf("Cannot create folder: %s", errCreate.Error())
	}

	FileName := fmt.Sprintf("%s/%d", Path, time.Now().UnixNano())
	part, err := writer.CreateFormFile(Path, FileName)
	if err != nil {
		log.Printf("Cannot create because file to big, %s", err.Error())
	}

	_, err = io.Copy(part, bytes.NewReader(buf))
	if err != nil {
		log.Printf("Cant coppy files, because: %s", err.Error())
	}

	return FileName
}