package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
HTTP is stateless between client and server, it means server not record any data to remember the every client
request in order to increase the scalabiliy of the server/

So how to remember a client request, in this case when user is login then not to ask login again for the next request.

The browser will save a cookie from the server, and every time a client make a request, a cookie will send automatically to the server.

*/

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// make a cookie with http.Cookie
	cookie := new(http.Cookie)
	cookie.Name = "X-Custom-Cookie"
	cookie.Value = r.URL.Query().Get("name") // set cookie value to name from query parameter

	// set cookie to writer
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/?name=satrya", nil)
	recorder := httptest.NewRecorder()

	SetCookieHandler(recorder, request)

	// get cookies
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("%s : %s \n", cookie.Name, cookie.Value)
	}
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Custom-Cookie-2")
	if err != nil {
		fmt.Fprint(w, "No cookie")
	} else {
		fmt.Fprintf(w, "Hello this is key from cookie %s with value %s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Custom-Cookie-2"
	cookie.Value = "CookieValue"

	// add cookie
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()
	GetCookieHandler(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
