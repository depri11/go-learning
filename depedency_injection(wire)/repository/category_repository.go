package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/depri11/go-learning/restful/helper"
	"github.com/depri11/go-learning/restful/model/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name) VALUES ($1)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.RowsAffected()
	helper.PanicIfError(err)

	category.ID = int(id)

	return category
}

func (r *repository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "UPDATE category SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.ID)
	helper.PanicIfError(err)

	return category
}

func (r *repository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "DELETE FROM category WHERE id = $1"
	_, err := tx.ExecContext(ctx, sql, category.ID)
	helper.PanicIfError(err)

	return category
}

func (r *repository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "SELECT id,name FROM category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (r *repository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id,name FROM category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
