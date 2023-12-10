package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func ReadFormFile(Path string, w http.ResponseWriter, r http.Request) string {
	const MAX_UPLOAD_FILE_SIZE = 100 * 1024 * 1024

	if err := r.ParseMultipartForm(MAX_UPLOAD_FILE_SIZE); err != nil {
		log.Printf("Cannot upload because file to big, %s", err.Error())
	}
	forms := r.MultipartForm.File["file"]

	var filesName string
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
		_, err = file.Read(buff)
		if err != nil {
			log.Printf("Cannot read File because, %s", err.Error())

		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Printf("Cannot read File Seek because, %s", err.Error())
		}

		errCreate := os.MkdirAll(Path, os.ModePerm)
		if errCreate != nil {
			log.Printf("Cannot create folder: %s", errCreate.Error())
		}

		fileName := fmt.Sprintf("%s/%d-%s", Path, time.Now().Day(), fileHeader.Filename)
		out, errs := os.Create(fileName)
		if errs != nil {
			log.Printf("Cannot create because file to big, %s", err.Error())
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Printf("Cant coppy files, because: %s", err.Error())
		}

		filesName = fileName
	}

	return filesName
}
