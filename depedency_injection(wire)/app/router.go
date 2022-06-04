package app

import (
	"github.com/depri11/go-learning/restful/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.Controller) *httprouter.Router {
	r := httprouter.New()

	r.GET("/api/categories", categoryController.FindAll)
	r.GET("/api/categories/:categoryId", categoryController.FindById)
	r.POST("/api/categories", categoryController.Create)
	r.PUT("/api/categories/:categoryId", categoryController.Update)
	r.DELETE("/api/categories/:categoryId", categoryController.Delete)

	return r
}
