package storage

import (
	"fmt"
	"cookbook/internal/entity"
	"cookbook/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgreSQLStorage struct {
	DB *sql.DB
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

	return &PostgreSQLStorage{DB: db}, nil
}


func (s *PostgreSQLStorage) GetAllDishes() ([]entity.DishOutput, error) {
	rows, err := s.DB.Query("SELECT name, description, time FROM schema_name.dish")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []entity.DishOutput
	for rows.Next() {
		var name, description string
		var time int

		err := rows.Scan(&name, &description, &time)
		if err != nil {
			return nil, err
		}

		dish := entity.DishOutput{
			Name: 		 name,
			Description: description,
			Time: 		 time,
		}

		dishes = append(dishes, dish)

	}

	return dishes, nil
}


func (s *PostgreSQLStorage) GetDishInfo(dishName string) (entity.DishInfo, error) {
    row := s.DB.QueryRow("SELECT id, name, description, recipe, time FROM schema_name.dish WHERE name = $1", dishName)

    var nameDish, description, recipe string
    var id, time int

    err := row.Scan(&id, &nameDish, &description, &recipe, &time)
    if err != nil {
        return entity.DishInfo{}, err
    }

    ingredientRows, err := s.DB.Query("SELECT i.name, i.measure_unit, di.quantity FROM schema_name.ingredient i JOIN schema_name.dish_ingredient di ON i.id = di.ingredient_id WHERE di.dish_id = $1", id)
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
		Name: 		 dishName,
		Description: description,
		Recipe: 	 recipe,
		Time: 		 time,
		Ingredients: ingredients,
	}

    return dishInfo, nil
}
