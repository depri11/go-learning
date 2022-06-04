package app

import (
	"database/sql"
	"time"

	"github.com/depri11/go-learning/restful/helper"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123 dbname=pzn_restful sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
