package Web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHelloWIthQueryHandler(writer http.ResponseWriter, r *http.Request) {
	// get query parameter
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintln(writer, "Hello anonymous")
	} else {
		fmt.Fprintln(writer, "Hello "+name)
	}
}

func TestSayHelloHandler(t *testing.T) {
	queryName := "Satrya"
	request := httptest.NewRequest("GET", "http://localhost:8081/hello?name="+queryName, nil)
	recorder := httptest.NewRecorder()

	SayHelloWIthQueryHandler(recorder, request)

	// result
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	expectedResult := "Hello " + queryName + "\n"
	assert.Equal(t, expectedResult, string(body))
}

func MultipleQueryParameterHandler(writer http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestMultipleQueryParams(t *testing.T) {
	firstName := "Satrya"
	lastName := "Budi"
	requestUrl := "http://localhost:8081/hello?first_name=" + firstName + "&last_name=" + lastName
	request := httptest.NewRequest("GET", requestUrl, nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameterHandler(recorder, request)
	//result
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	expectedResult := firstName + " " + lastName
	assert.Equal(t, expectedResult, string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["same_key"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/hello?same_key=valueA&same_key=valueB&same_key=valueC", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
