package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Root path requested\n")
	io.WriteString(w, "Status code 200 (ok)\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello path requested\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe("127.0.0.1:3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed succsefully\n")
	} else if err != nil {
		fmt.Printf("server is not started %s\n", err)
		os.Exit(1)
	}
}
