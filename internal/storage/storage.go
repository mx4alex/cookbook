package storage

import (
	"cookbook/internal/entity"
	"database/sql"
)

type Dish interface {
	GetAllDishes() ([]entity.Dish, error)
	GetDishInfo(int) (entity.Dish, error)
	AddDish(*entity.Dish) error
	UpdateDish(int, *entity.Dish) error
	DeleteDish(int) error
}

type Cousine interface {
	GetCousines() ([]entity.Cousine, error)
	AddCousine(*entity.Cousine) (int, error)
	UpdateCousine(int, *entity.Cousine) error
	DeleteCousine(int) error
}

type Category interface {

}

type Storage struct {
	Dish
	Cousine
	Category
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Dish: 		NewDishPostgres(db),
		Cousine:    NewCousinePostgres(db),
		Category:   NewCategoryPostgres(db),
	}
}