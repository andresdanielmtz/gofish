package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading file...!")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	} // 10 MB limit

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file) // defer means that this will be executed at the end of the function

	// Before uploading, be sure to check the folder exists
	err = os.MkdirAll("images", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("File Size: %+v bytes\n", handler.Size)
	tempFile, err := os.CreateTemp("images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "Uploaded: %s", tempFile.Name())
	if err != nil {
		return
	}
}
