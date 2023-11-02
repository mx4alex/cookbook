package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type CategoryPostgres struct {
	db *sql.DB
}

func NewCategoryPostgres(db *sql.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}