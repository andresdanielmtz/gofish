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
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
		}
	}(tempFile)

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

func GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("images")
	if err != nil {
		http.Error(w, "Could not read directory", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write([]byte(`<a href="/">Home</a><br>`))
	if err != nil {
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			link := fmt.Sprintf(`<a href="/images/%s">%s</a> <a href="/download/%s">download </a><br>`, file.Name(), file.Name(), file.Name())
			_, err := w.Write([]byte(link))
			if err != nil {
				return
			}
		}
	}
}

func GetFileByIDHandler(w http.ResponseWriter, r *http.Request) {
	nameFile := r.URL.Path[len("/images/"):]

	if nameFile == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}

	filePath, err := getFileById(nameFile)
	fmt.Println("Serving file:", filePath, "Error:", err) // Debug print
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	w.Header().Set("Content-Type", "image/png")
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Could not write file", http.StatusInternalServerError)
		return
	}
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	nameFile := r.URL.Path[len("/download/"):]

	if nameFile == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}
	filePath, err := getFileById(nameFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	w.Header().Set("Content-Disposition", "attachment; filename="+nameFile)
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Could not write file", http.StatusInternalServerError)
		return
	}
}
