package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Root path function added
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Root path requested\n")
	fmt.Fprint(w, "<h1>This is my website</h1>")
}

// /hello path function added
func getHello(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("/hello path requested\n")
	fmt.Fprint(w, "<h1>This is my website with /hello path selected</h1>")
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
		fmt.Println("Closing, thank you!")
		os.Exit(1)
	}
}
