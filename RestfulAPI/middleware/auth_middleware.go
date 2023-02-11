package middleware

import (
	"github.com/noczero/Golang-Hands-On/RestfulAPI/exception"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/helper"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := "SECRETAPIKEYEXAMPLE"
	if apiKey == request.Header.Get("X-API-Key") {
		//request header X-API-Key match with our api key then forward the middleware handler
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// not match UNAUTHORIZATION
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		unauthorizedError := exception.NewUnAutohorizedError("API key is invalid")
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   unauthorizedError,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
