package main

import (
	"net/http"

	"github.com/depri11/go-learning/restful/app"
	"github.com/depri11/go-learning/restful/controller"
	"github.com/depri11/go-learning/restful/helper"
	"github.com/depri11/go-learning/restful/middleware"
	"github.com/depri11/go-learning/restful/repository"
	"github.com/depri11/go-learning/restful/service"
	"github.com/depri11/go-learning/restful/exception"
	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()

	validate := validator.New()
	categoryRepository := repository.NewRepository()
	categoryService := service.NewService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	r := app.NewRouter(categoryController)

	r.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(r),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
