package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
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
	GetCategories() ([]entity.Category, error)
	AddCategory(*entity.Category) (int, error)
	UpdateCategory(int, *entity.Category) error
	DeleteCategory(int) error
}

type Service struct {
	Dish
	Cousine
	Category
}

func NewService(st *storage.Storage) *Service {
	return &Service{
		Dish: 		NewDishService(st.Dish),
		Cousine:    NewCousineService(st.Cousine),
		Category:   NewCategoryService(st.Category),
	}
}