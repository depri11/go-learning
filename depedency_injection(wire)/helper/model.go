package helper

import (
	"github.com/depri11/go-learning/restful/model/domain"
	"github.com/depri11/go-learning/restful/model/web"
)

func ToCreateResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCreateResponse(category))
	}
	return categoryResponses
}
