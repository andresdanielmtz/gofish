package main

import (
	"errors"
	"fmt"
	"gofish/api"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", api.GetRoot)
	http.HandleFunc("/hello", api.GetHello)
	http.HandleFunc("/upload", api.UploadFileHandler)
	http.HandleFunc("/images", api.GetFilesHandler)
	http.HandleFunc("/images/", api.GetFileByIDHandler)
	http.HandleFunc("/download", api.DownloadFileHandler)

	fmt.Println("Starting HTTP server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("HTTP server closed")
	} else if err != nil {
		fmt.Printf("Error in ListenAndServe: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Hello, World!")
}
