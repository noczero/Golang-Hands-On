package Web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCodeHandler(w http.ResponseWriter, r *http.Request) {
	// if there's query parameter then return the data with response code 200,
	// if empty return response code 400 means BAD REQUEST
	name := r.URL.Query().Get("name")

	if name == "" {
		// change status code to 400
		w.WriteHeader(400)
		fmt.Fprint(w, "Name is empty, BAD REQUEST")
	} else {
		// change status code to 200
		w.WriteHeader(200)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponsCodeSuccess(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/response_code?name=satrya", nil)
	recorder := httptest.NewRecorder()

	ResponseCodeHandler(recorder, request)
	assert.Equal(t, 200, recorder.Code)

	response := recorder.Result()
	fmt.Println(response.StatusCode) // 200
	fmt.Println(response.Status)     // 200 OK

	// body result
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TestResponseCodeBadRequest(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/response_code?", nil)
	recorder := httptest.NewRecorder()

	ResponseCodeHandler(recorder, request)
	assert.Equal(t, 400, recorder.Code)

	response := recorder.Result()
	fmt.Println(response.StatusCode) // 200
	fmt.Println(response.Status)     // 400 Bad Request

	// body result
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
