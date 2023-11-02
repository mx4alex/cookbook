package storage

import (
	"fmt"
	"cookbook/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

func New(cfg config.PostgresConfig) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}