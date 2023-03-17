package Web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	// Handler is used to receive the request from client and write the response to clinet
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello world")
	}

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	// ServerMux is a handler to support multiple endpoints.
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World~")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	server := http.Server{Addr: "localhost:8081", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServerMuxURLPattern(t *testing.T) {

	mux := http.NewServeMux()

	// By default, /images/ will match to the any URL after /images/*, like /images/img1, /images/img2
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images")
	})

	// If there's a handler function with specific after /images/ it will match current handlers.
	mux.HandleFunc("/images/thumbnails", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnails")
	})

	server := http.Server{Addr: "localhost:8081", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)     // this will print GET
		fmt.Fprintln(writer, request.RequestURI) // this will print path URI
	}

	server := http.Server{Addr: "localhost:8081", Handler: handler}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
