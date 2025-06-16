package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Serves main HTML file for the root path. (!!)
func GetRoot(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("static/index.html")
	if err != nil {
		fmt.Println("Error opening index.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
