package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Root handler for the web server.
// These are mostly used for serving static files or HTML content.
// /hello is used to test the server's response.

// !! Serves main HTML file for the root path.

func GetRoot(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("static/index.html")
	if err != nil {
		fmt.Println("Error opening index.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		return
	}
}
