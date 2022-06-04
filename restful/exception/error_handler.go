package exception

import (
	"net/http"

	"github.com/depri11/go-learning/restful/helper"
	"github.com/depri11/go-learning/restful/model/web"
	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if exception, ok := err.(NotFoundError); ok {
		w.Header().Set("Content-Type", "application/")
		w.WriteHeader(http.StatusInternalServerError)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
