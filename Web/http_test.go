package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Worlds!")
}

func TestHelloWorldHandler(t *testing.T) {
	// using htttest without running the server.ListenNServer/

	// make request, define method and dummy path.
	request := httptest.NewRequest("GET", "http://localhost:8081/hello", nil)

	// make writer as recorder that accepts response
	recorder := httptest.NewRecorder()

	HelloWorldHandler(recorder, request)

	// check response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
