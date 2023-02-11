package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/controller"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/exception"
)

func NewRouter(controller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	// list of endpoints
	router.GET("/api/categories", controller.FindAll)
	router.POST("/api/categories", controller.Create)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.PUT("/api/categories/:categoryId", controller.Update)
	router.DELETE("/api/categories/:categoryId", controller.Delete)

	// add exception
	router.PanicHandler = exception.ErrorHandler

	return router
}
