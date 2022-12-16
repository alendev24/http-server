package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Root path function added
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Root path requested\n")
	io.WriteString(w, "Status code 200 (ok)\n")
}

// /hello path function added
func getHello(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	ctx := r.Context()

	fmt.Printf("%s: got/hello request\n", ctx.Value(keyServerAddr))
=======
>>>>>>> parent of 1607448 (added two servers)
	fmt.Printf("/hello path requested\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

// Program main entry point
func main() {
<<<<<<< HEAD
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error occured for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error occured for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
=======
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
>>>>>>> parent of 1607448 (added two servers)
}
