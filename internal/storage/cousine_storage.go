package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type CousinePostgres struct {
	db *sql.DB
}

func NewCousinePostgres(db *sql.DB) *CousinePostgres {
	return &CousinePostgres{db: db}
}