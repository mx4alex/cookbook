package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
	"context"
)

type Dish interface {
	GetAllDishes(context.Context) ([]entity.Dish, error)
	GetDishInfo(context.Context, int) (entity.Dish, error)
	AddDish(context.Context, *entity.Dish) (int, error)
	UpdateDish(context.Context, int, *entity.Dish) error
	DeleteDish(context.Context, int) error
	
	GetDishCousine(context.Context, int) ([]entity.Dish, error)
	GetDishCategory(context.Context, int) ([]entity.Dish, error)
	GetDishCousineCategory(context.Context, int, int) ([]entity.Dish, error)
}

type Cousine interface {
	GetCousines(context.Context) ([]entity.Cousine, error)
	AddCousine(context.Context, *entity.Cousine) (int, error)
	UpdateCousine(context.Context, int, *entity.Cousine) error
	DeleteCousine(context.Context, int) error
}

type Category interface {
	GetCategories(context.Context) ([]entity.Category, error)
	AddCategory(context.Context, *entity.Category) (int, error)
	UpdateCategory(context.Context, int, *entity.Category) error
	DeleteCategory(context.Context, int) error
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