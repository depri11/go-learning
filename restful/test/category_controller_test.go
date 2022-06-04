package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/depri11/go-learning/restful/app"
	"github.com/depri11/go-learning/restful/controller"
	"github.com/depri11/go-learning/restful/helper"
	"github.com/depri11/go-learning/restful/middleware"
	"github.com/depri11/go-learning/restful/model/domain"
	"github.com/depri11/go-learning/restful/repository"
	"github.com/depri11/go-learning/restful/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123 dbname=pzn_restful_test sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewRepository()
	categoryService := service.NewService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	r := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(r)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"test"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories"+strconv.Itoa(category.ID), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.ID, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}

func TestListCategoriesSuccess(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
