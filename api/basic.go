package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetRoot is the main handler for the / route. It delivers the index.html from the static folder.
func GetRoot(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("static/index.html") // ? It seems file path always comes from the root instead from a relative directory. Sounds good (y)
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

// GetHello is a test route. Used only for server testing purposes.
func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		return
	}
}
