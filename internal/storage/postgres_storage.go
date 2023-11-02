package storage

import (
	"fmt"
	"cookbook/internal/entity"
	"cookbook/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgreSQLStorage struct {
	db *sql.DB
}

func New(cfg config.PostgresConfig) (*PostgreSQLStorage, error) {
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

	return &PostgreSQLStorage{db: db}, nil
}


func (s *PostgreSQLStorage) GetAllDishes() ([]entity.DishOutput, error) {
	rows, err := s.db.Query("SELECT id, name, description, time FROM test.dish")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []entity.DishOutput
	for rows.Next() {
		var name, description string
		var dishID, time int

		err := rows.Scan(&dishID, &name, &description, &time)
		if err != nil {
			return nil, err
		}

		dish := entity.DishOutput{
			ID:			 dishID,
			Name: 		 name,
			Description: description,
			Time: 		 time,
		}

		dishes = append(dishes, dish)

	}

	return dishes, nil
}


func (s *PostgreSQLStorage) GetDishInfo(dishID int) (entity.DishInfo, error) {
    row := s.db.QueryRow("SELECT id, name, description, recipe, time FROM test.dish WHERE id = $1", dishID)

    var dishName, description, recipe string
    var id, time int

    err := row.Scan(&id, &dishName, &description, &recipe, &time)
    if err != nil {
        return entity.DishInfo{}, err
    }

    ingredientRows, err := s.db.Query("SELECT i.name, i.measure_unit, di.quantity FROM test.ingredient i JOIN test.dish_ingredient di ON i.id = di.ingredient_id WHERE di.dish_id = $1", id)
    if err != nil {
        return entity.DishInfo{}, err
    }
    defer ingredientRows.Close()

    var ingredients []entity.Ingredient

    for ingredientRows.Next() {
        var ingredientName, measureUnit string
        var quantity int

        err := ingredientRows.Scan(&ingredientName, &measureUnit, &quantity)
        if err != nil {
            return entity.DishInfo{}, err
        }

		ingredient := entity.Ingredient{
			Name:   	 ingredientName,
			MeasureUnit: measureUnit,
			Quantity: 	 quantity,
		}

        ingredients = append(ingredients, ingredient)
    }

	dishInfo := entity.DishInfo{
		ID:			 dishID,
		Name: 		 dishName,
		Description: description,
		Recipe: 	 recipe,
		Time: 		 time,
		Ingredients: ingredients,
	}

    return dishInfo, nil
}

func (s *PostgreSQLStorage) AddDish(dishInput *entity.DishInput) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO test.dish (name, category_id, cousine_id, description, recipe, time) VALUES ($1, $2, $3, $4, $5, $6)",
		dishInput.Name, dishInput.CategoryId, dishInput.CousineId, dishInput.Description, dishInput.Recipe, dishInput.Time)
	if err != nil {
		return err
	}

	var dishID int
	err = tx.QueryRow("SELECT id FROM test.dish WHERE name = $1", dishInput.Name).Scan(&dishID)
	if err != nil {
		return err
	}

	for _, ingredient := range dishInput.Ingredients {

		var exists bool
		err := tx.QueryRow("SELECT EXISTS(SELECT * FROM test.ingredient WHERE name = $1)", ingredient.Name).Scan(&exists)
		if err != nil {
			return err
		}
		
		if !exists {
			_, err = tx.Exec("INSERT INTO test.ingredient (name, measure_unit, protein, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5)",
				ingredient.Name, ingredient.MeasureUnit, 0, 0, 0)
			if err != nil {
				return err
			}
		} 

		var ingredientID int
		err = tx.QueryRow("SELECT id FROM test.ingredient WHERE name = $1", ingredient.Name).Scan(&ingredientID)
		if err != nil {
			return err
		}
		_, err = tx.Exec("INSERT INTO test.dish_ingredient (dish_id, ingredient_id, quantity) VALUES ($1, $2, $3)",
			dishID, ingredientID, ingredient.Quantity)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreSQLStorage) UpdateDish(dishID int, dishInput *entity.DishInput) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE test.dish SET name = $2, category_id = $3, cousine_id = $4, description = $5, recipe = $6, time = $7 WHERE id = $1",
		dishID, dishInput.Name, dishInput.CategoryId, dishInput.CousineId, dishInput.Description, dishInput.Recipe, dishInput.Time)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM test.dish_ingredient WHERE dish_id = $1", dishID)
	if err != nil {
		return err
	}

	for _, ingredient := range dishInput.Ingredients {

		var exists bool
		err := tx.QueryRow("SELECT EXISTS(SELECT * FROM test.ingredient WHERE name = $1)", ingredient.Name).Scan(&exists)
		if err != nil {
			return err
		}
		
		if !exists {
			_, err = tx.Exec("INSERT INTO test.ingredient (name, measure_unit, protein, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5)",
				ingredient.Name, ingredient.MeasureUnit, 0, 0, 0)
			if err != nil {
				return err
			}
		} 

		var ingredientID int
		err = tx.QueryRow("SELECT id FROM test.ingredient WHERE name = $1", ingredient.Name).Scan(&ingredientID)
		if err != nil {
			return err
		}
		_, err = tx.Exec("INSERT INTO test.dish_ingredient (dish_id, ingredient_id, quantity) VALUES ($1, $2, $3)",
			dishID, ingredientID, ingredient.Quantity)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreSQLStorage) DeleteDish(dishID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	_, err = tx.Exec("DELETE FROM test.dish_ingredient WHERE dish_id = $1", dishID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM test.dish WHERE id = $1", dishID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

