package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
HTTP Header is additional information that send by client to server, or reversely.
It includes in HTTP request and HTTP response.
It's not case-sensitive unlike the query parameter that case-sensitive.
it will accept "content-type" or "Content-Type"
*/

func RequestHeaderHandler(writer http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8081/request", nil)
	request.Header.Add("content-type", "application/json")
	recoreder := httptest.NewRecorder()

	RequestHeaderHandler(recoreder, request)

	response := recoreder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeaderHandler(writer http.ResponseWriter, r *http.Request) {
	// add header to response
	writer.Header().Add("X-Dummy-Header", "custom header value")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8081", nil)
	recoreder := httptest.NewRecorder()

	ResponseHeaderHandler(recoreder, request)

	//response := recoreder.Result()
	//responseHeader := response.Header.Get("X-Dummy-Header")
	//
	responseHeader := recoreder.Header().Get("X-Dummy-Header")

	fmt.Println(responseHeader)
}
