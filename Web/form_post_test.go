package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
In RESTFul API, POST method is used via body HTTP request.
To get the body request we do ParseForm and get the data from PostForm.
*/

func FormPostHandler(w http.ResponseWriter, r *http.Request) {
	// parse post form
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// get the data
	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	// result
	fmt.Fprintf(w, "%s %s", firstName, lastName)

}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Satrya&lastName=Budi")

	// make a request
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPostHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
