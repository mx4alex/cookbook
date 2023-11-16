package storage

import (
	"cookbook/internal/entity"
	"database/sql"
	_ "github.com/lib/pq"
	"context"
)

type DishPostgres struct {
	db *sql.DB
}

func NewDishPostgres(db *sql.DB) *DishPostgres {
	return &DishPostgres{db: db}
}

func (s *DishPostgres) GetAllDishes(ctx context.Context) ([]entity.Dish, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, description, time FROM test.dish")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []entity.Dish
	var name, description string
	var dishID, time int

	for rows.Next() {
		err := rows.Scan(&dishID, &name, &description, &time)
		if err != nil {
			return nil, err
		}

		dish := entity.Dish{
			ID:			 dishID,
			Name: 		 name,
			Description: description,
			Time: 		 time,
		}

		dishes = append(dishes, dish)

	}

	return dishes, nil
}


func (s *DishPostgres) GetDishInfo(ctx context.Context, dishID int) (entity.Dish, error) {
    row := s.db.QueryRowContext(ctx, "SELECT id, name, cousine_id, category_id, description, recipe, time FROM test.dish WHERE id = $1", dishID)

    var dishName, description, recipe string
    var id, cousineID, categoryID, time int

    err := row.Scan(&id, &dishName, &cousineID, &categoryID, &description, &recipe, &time)
    if err != nil {
        return entity.Dish{}, err
    }

    ingredientRows, err := s.db.QueryContext(ctx, `SELECT i.name, i.measure_unit, i.protein, i.fats, i.carbohydrates, di.quantity FROM test.ingredient i 
													JOIN test.dish_ingredient di ON i.id = di.ingredient_id WHERE di.dish_id = $1`, id)
    if err != nil {
        return entity.Dish{}, err
    }
    defer ingredientRows.Close()

    var ingredients []entity.Ingredient
	var ingredientName, measureUnit string
	var protein, fats, carbohydrates, quantity int

    for ingredientRows.Next() {
        err := ingredientRows.Scan(&ingredientName, &measureUnit, &protein, &fats, &carbohydrates, &quantity)
        if err != nil {
            return entity.Dish{}, err
        }

		ingredient := entity.Ingredient{
			Name:   	 	ingredientName,
			MeasureUnit: 	measureUnit,
			Quantity: 	 	quantity,
			Protein: 	 	protein,
			Fats:        	fats,
			Carbohydrates: 	carbohydrates,
		}

        ingredients = append(ingredients, ingredient)
    }

	dishInfo := entity.Dish{
		ID:			 dishID,
		Name: 		 dishName,
		CousineID:   cousineID,
		CategoryID:  categoryID,
		Description: description,
		Recipe: 	 recipe,
		Time: 		 time,
		Ingredients: ingredients,
	}

    return dishInfo, nil
}

func (s *DishPostgres) AddDish(ctx context.Context, dishInput *entity.Dish) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "INSERT INTO test.dish (name, category_id, cousine_id, description, recipe, time) VALUES ($1, $2, $3, $4, $5, $6)",
		dishInput.Name, dishInput.CategoryID, dishInput.CousineID, dishInput.Description, dishInput.Recipe, dishInput.Time)
	if err != nil {
		return 0, err
	}

	var dishID int
	err = tx.QueryRowContext(ctx, "SELECT id FROM test.dish WHERE name = $1", dishInput.Name).Scan(&dishID)
	if err != nil {
		return 0, err
	}

	for _, ingredient := range dishInput.Ingredients {

		var exists bool
		err := tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT * FROM test.ingredient WHERE name = $1)", ingredient.Name).Scan(&exists)
		if err != nil {
			return 0, err
		}
		
		if !exists {
			_, err = tx.ExecContext(ctx, "INSERT INTO test.ingredient (name, measure_unit, protein, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5)",
				ingredient.Name, ingredient.MeasureUnit, 0, 0, 0)
			if err != nil {
				return 0, err
			}
		} 

		var ingredientID int
		err = tx.QueryRowContext(ctx, "SELECT id FROM test.ingredient WHERE name = $1", ingredient.Name).Scan(&ingredientID)
		if err != nil {
			return 0, err
		}
		_, err = tx.ExecContext(ctx, "INSERT INTO test.dish_ingredient (dish_id, ingredient_id, quantity) VALUES ($1, $2, $3)",
			dishID, ingredientID, ingredient.Quantity)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return dishID, nil
}

func (s *DishPostgres) UpdateDish(ctx context.Context, dishID int, dishInput *entity.Dish) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE test.dish SET name = $2, category_id = $3, cousine_id = $4, description = $5, recipe = $6, time = $7 WHERE id = $1",
		dishID, dishInput.Name, dishInput.CategoryID, dishInput.CousineID, dishInput.Description, dishInput.Recipe, dishInput.Time)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM test.dish_ingredient WHERE dish_id = $1", dishID)
	if err != nil {
		return err
	}

	var exists bool
	var count int

	for _, ingredient := range dishInput.Ingredients {

		err := tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT * FROM test.ingredient WHERE name = $1)", ingredient.Name).Scan(&exists)
		if err != nil {
			return err
		}
		
		if !exists {
			_, err = tx.ExecContext(ctx, "INSERT INTO test.ingredient (name, measure_unit, protein, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5)",
				ingredient.Name, ingredient.MeasureUnit, ingredient.Protein, ingredient.Fats, ingredient.Carbohydrates)
			if err != nil {
				return err
			}
		} else {
			err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM test.ingredient i JOIN test.dish_ingredient di ON i.id = di.ingredient_id WHERE i.name = $1", ingredient.Name).Scan(&count)
			if err != nil {
				return err
			}

			if count == 0 {
				_, err = tx.ExecContext(ctx, "UPDATE test.ingredient SET measure_unit = $2, protein = $3, fats = $4, carbohydrates = $5 WHERE name = $1",
					ingredient.Name, ingredient.MeasureUnit, ingredient.Protein, ingredient.Fats, ingredient.Carbohydrates)
				if err != nil {
					return err
				}
			}
		}

		var ingredientID int
		err = tx.QueryRowContext(ctx, "SELECT id FROM test.ingredient WHERE name = $1", ingredient.Name).Scan(&ingredientID)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, "INSERT INTO test.dish_ingredient (dish_id, ingredient_id, quantity) VALUES ($1, $2, $3)",
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

func (s *DishPostgres) DeleteDish(ctx context.Context, dishID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	_, err = tx.ExecContext(ctx, "DELETE FROM test.dish_ingredient WHERE dish_id = $1", dishID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM test.dish WHERE id = $1", dishID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *DishPostgres) GetDishCousine(ctx context.Context, cousineID int) ([]entity.Dish, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, description, time FROM test.dish WHERE cousine_id = $1", cousineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []entity.Dish
	var name, description string
	var dishID, time int

	for rows.Next() {
		err := rows.Scan(&dishID, &name, &description, &time)
		if err != nil {
			return nil, err
		}

		dish := entity.Dish{
			ID:			 dishID,
			Name: 		 name,
			Description: description,
			Time: 		 time,
		}

		dishes = append(dishes, dish)

	}

	return dishes, nil
}

func (s *DishPostgres) GetDishCategory(ctx context.Context, categoryID int) ([]entity.Dish, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, description, time FROM test.dish WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []entity.Dish
	var name, description string
	var dishID, time int

	for rows.Next() {
		err := rows.Scan(&dishID, &name, &description, &time)
		if err != nil {
			return nil, err
		}

		dish := entity.Dish{
			ID:			 dishID,
			Name: 		 name,
			Description: description,
			Time: 		 time,
		}

		dishes = append(dishes, dish)

	}

	return dishes, nil
}

func (s *DishPostgres) GetDishCousineCategory(ctx context.Context, cousineID, categoryID int) ([]entity.Dish, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, description, time FROM test.dish WHERE cousine_id = $1 AND category_id = $2", cousineID, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []entity.Dish
	var name, description string
	var dishID, time int

	for rows.Next() {
		err := rows.Scan(&dishID, &name, &description, &time)
		if err != nil {
			return nil, err
		}

		dish := entity.Dish{
			ID:			 dishID,
			Name: 		 name,
			Description: description,
			Time: 		 time,
		}

		dishes = append(dishes, dish)

	}

	return dishes, nil
}

func (s *DishPostgres) GetDishSearch(ctx context.Context, words []string) ([]entity.Dish, error) {
	if len(words) == 0 {
		return nil, nil
	}

	var dishes []entity.Dish
	var name, description string
	var dishID, time int
	setID := make(map[int]bool)

	for _, word := range words {
	
		rowsDish, err := s.db.QueryContext(ctx, "SELECT id, name, description, time FROM test.dish WHERE name LIKE $1", word + string('%'))
		if err != nil {
			return nil, err
		}
		defer rowsDish.Close()

		for rowsDish.Next() {
			err := rowsDish.Scan(&dishID, &name, &description, &time)
			if err != nil {
				return nil, err
			}

			_, ok := setID[dishID]
			if !ok {
				setID[dishID] = true

				dish := entity.Dish{
					ID:			 dishID,
					Name: 		 name,
					Description: description,
					Time: 		 time,
				}
		
				dishes = append(dishes, dish)
			}

		}

		rowsDishIngredient, err := s.db.QueryContext(ctx, `SELECT d.id, d.name, d.description, d.time FROM test.dish d JOIN test.dish_ingredient di ON d.id = di.dish_id 
															JOIN test.ingredient i ON di.ingredient_id = i.id WHERE i.name LIKE $1`, word + string('%'))
		if err != nil {
			return nil, err
		}
		defer rowsDishIngredient.Close()
		
		for rowsDishIngredient.Next() {
			err := rowsDishIngredient.Scan(&dishID, &name, &description, &time)
			if err != nil {
				return nil, err
			}

			_, ok := setID[dishID]
			if !ok {
				setID[dishID] = true

				dish := entity.Dish{
					ID:			 dishID,
					Name: 		 name,
					Description: description,
					Time: 		 time,
				}
		
				dishes = append(dishes, dish)
			}
		}
	}

	return dishes, nil
}