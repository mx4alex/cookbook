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

}

type Category interface {

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