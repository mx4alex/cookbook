package storage

import (
	"cookbook/internal/entity"
	"database/sql"
	_ "github.com/lib/pq"
)

type CousinePostgres struct {
	db *sql.DB
}

func NewCousinePostgres(db *sql.DB) *CousinePostgres {
	return &CousinePostgres{db: db}
}

func (s *CousinePostgres) GetCousines() ([]entity.Cousine, error) {
	rows, err := s.db.Query("SELECT id, name, description FROM test.cousine")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cousines []entity.Cousine
	for rows.Next() {
		var name, description string
		var cousineID int

		err := rows.Scan(&cousineID, &name, &description)
		if err != nil {
			return nil, err
		}

		cousine := entity.Cousine {
			ID:			 cousineID,
			Name: 		 name,
			Description: description,
		}

		cousines = append(cousines, cousine)

	}

	return cousines, nil
}

func (s *CousinePostgres) AddCousine(cousine *entity.Cousine) (int, error) {
	_, err := s.db.Exec("INSERT INTO test.cousine (name, description) VALUES ($1, $2)", cousine.Name, cousine.Description)
	if err != nil {
		return 0, err
	}

	var cousineID int
	err = s.db.QueryRow("SELECT id FROM test.cousine WHERE name = $1", cousine.Name).Scan(&cousineID)
	if err != nil {
		return 0, err
	}

	return cousineID, nil
}

func (s *CousinePostgres) UpdateCousine(cousineID int, cousine *entity.Cousine) error {
	_, err := s.db.Exec("UPDATE test.cousine SET name = $2, description = $3 WHERE id = $1", cousineID, cousine.Name, cousine.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s *CousinePostgres) DeleteCousine(cousineID int) error {
	_, err := s.db.Exec("DELETE FROM test.cousine WHERE id = $1", cousineID)
	if err != nil {
		return err
	}

	return nil
}
