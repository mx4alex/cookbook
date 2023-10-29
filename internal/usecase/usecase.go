package usecase

import (
	"cookbook/internal/entity"
)

type DishStorage interface {
	GetAllDishes() ([]entity.DishOutput, error)
	GetDishInfo(name string) (entity.DishInfo, error)
	AddDish(*entity.DishInput) error
}

type CookInteractor struct {
	dishStorage DishStorage
}

func NewTaskInteractor(dishStorage DishStorage) *CookInteractor {
	return &CookInteractor{dishStorage: dishStorage}
}

func (c *CookInteractor) GetAllDishes() ([]entity.DishOutput, error) {
	return c.dishStorage.GetAllDishes()
}

func (c *CookInteractor) GetDishInfo(name string) (entity.DishInfo, error) {
	return c.dishStorage.GetDishInfo(name)
}

func (c *CookInteractor) AddDish(dish *entity.DishInput) error {
	return c.dishStorage.AddDish(dish)
}
