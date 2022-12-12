package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Root path function added
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Root path requested\n")
	io.WriteString(w, "Status code 200 (ok)\n")
}

// /hello path function added
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello path requested\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

// Program main entry point
func main() {
	// Registering function as handler for specific path, in this case getRoot and getHello
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	// Error handling in specific cases
	err := http.ListenAndServe("127.0.0.1:3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed succsefully\n")
	} else if err != nil {
		fmt.Printf("server is not started %s\n", err)
		os.Exit(1)
	}
}
