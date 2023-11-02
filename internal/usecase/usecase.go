package usecase

import (
	"cookbook/internal/entity"
)

type DishStorage interface {
	GetAllDishes() ([]entity.DishOutput, error)
	GetDishInfo(int) (entity.DishInfo, error)
	AddDish(*entity.DishInput) error
	UpdateDish(int, *entity.DishInput) error
	DeleteDish(int) error
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

func (c *CookInteractor) GetDishInfo(id int) (entity.DishInfo, error) {
	return c.dishStorage.GetDishInfo(id)
}

func (c *CookInteractor) AddDish(dish *entity.DishInput) error {
	return c.dishStorage.AddDish(dish)
}

func (c *CookInteractor) UpdateDish(id int, dish *entity.DishInput) error {
	return c.dishStorage.UpdateDish(id, dish)
}

func (c *CookInteractor) DeleteDish(id int) error {
	return c.dishStorage.DeleteDish(id)
}
