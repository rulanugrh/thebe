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

func ReadFormFile(path string, w http.ResponseWriter, r http.Request) (string, string) {
	const MAX_UPLOAD_FILE_SIZE = 100 * 1024 * 1024

	if err := r.ParseMultipartForm(MAX_UPLOAD_FILE_SIZE); err != nil {
		log.Printf("Cannot upload because file to big, %s", err.Error())
	}
	forms := r.MultipartForm.File["file"]

	var filesName string
	var boundary string
	for _, fileHeader := range forms {

		if fileHeader.Size > MAX_UPLOAD_FILE_SIZE {
			log.Printf("Cannot upload because file to big")
		}

		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("Cannot form File because, %s", err.Error())

		}
		defer file.Close()

		buff := make([]byte, 1024)
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		boundarys := writer.Boundary()

		_, err = file.Read(buff)
		if err != nil {
			log.Printf("Cannot read File because, %s", err.Error())

		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Printf("Cannot read File Seek because, %s", err.Error())
		}

		errCreate := os.MkdirAll(path, os.ModePerm)
		if errCreate != nil {
			log.Printf("Cannot create folder: %s", errCreate.Error())
		}

		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)
		folders := fmt.Sprintf("%s/%s", path, fileName)

		out, errs := os.Create(folders)
		if errs != nil {
			log.Printf("Cannot create because file to big, %s", err.Error())
		}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Printf("Cant coppy files, because: %s", err.Error())
		}

		filesName = fileName
		boundary = boundarys
	}

	return filesName, boundary
}
