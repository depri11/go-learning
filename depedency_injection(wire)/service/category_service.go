package service

import (
	"context"
	"database/sql"

	"github.com/depri11/go-learning/restful/exception"
	"github.com/depri11/go-learning/restful/helper"
	"github.com/depri11/go-learning/restful/model/domain"
	"github.com/depri11/go-learning/restful/model/web"
	"github.com/depri11/go-learning/restful/repository"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}

type service struct {
	repository repository.Repository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewService(repository repository.Repository, DB *sql.DB, validate *validator.Validate) Service {
	return &service{
		repository: repository,
		DB:         DB,
		Validate:   validate}
}

func (s *service) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = s.repository.Save(ctx, tx, category)

	return helper.ToCreateResponse(category)
}

func (s *service) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.repository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = s.repository.Update(ctx, tx, category)

	return helper.ToCreateResponse(category)
}

func (s *service) Delete(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.repository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.repository.Delete(ctx, tx, category)
}

func (s *service) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.repository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCreateResponse(category)
}

func (s *service) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := s.repository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
