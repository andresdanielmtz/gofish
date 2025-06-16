package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request")
	_, err := io.WriteString(w, "This is my website!\n")
	if err != nil {
		return
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("HTTP server closed")
	} else if err != nil {
		fmt.Printf("Error in ListenAndServe: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Hello, World!")
}
