package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/app"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/controller"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/helper"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/middleware"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/repository"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// tidy up the controller
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		//Handler: router,

		// add middleware as router
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
