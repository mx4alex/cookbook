package storage

import (
	"cookbook/internal/entity"
	"database/sql"
	_ "github.com/lib/pq"
)

type CategoryPostgres struct {
	db *sql.DB
}

func NewCategoryPostgres(db *sql.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (s *CategoryPostgres) GetCategories() ([]entity.Category, error) {
	rows, err := s.db.Query("SELECT id, name, description FROM test.category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var name, description string
		var categoryID int

		err := rows.Scan(&categoryID, &name, &description)
		if err != nil {
			return nil, err
		}

		category := entity.Category {
			ID:			 categoryID,
			Name: 		 name,
			Description: description,
		}

		categories = append(categories, category)

	}

	return categories, nil
}

func (s *CategoryPostgres) AddCategory(category *entity.Category) (int, error) {
	_, err := s.db.Exec("INSERT INTO test.category (name, description) VALUES ($1, $2)", category.Name, category.Description)
	if err != nil {
		return 0, err
	}

	var categoryID int
	err = s.db.QueryRow("SELECT id FROM test.category WHERE name = $1", category.Name).Scan(&categoryID)
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}

func (s *CategoryPostgres) UpdateCategory(categoryID int, category *entity.Category) error {
	_, err := s.db.Exec("UPDATE test.category SET name = $2, description = $3 WHERE id = $1", categoryID, category.Name, category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryPostgres) DeleteCategory(categoryID int) error {
	_, err := s.db.Exec("DELETE FROM test.category WHERE id = $1", categoryID)
	if err != nil {
		return err
	}

	return nil
}
