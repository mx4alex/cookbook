package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
)

type DishService struct {
	st storage.Dish
}

func NewDishService(st storage.Dish) *DishService {
	return &DishService{st: st}
}

func (s *DishService) GetAllDishes() ([]entity.Dish, error) {
	return s.st.GetAllDishes()
}

func (s *DishService) GetDishInfo(id int) (entity.Dish, error) {
	return s.st.GetDishInfo(id)
}

func (s *DishService) AddDish(dish *entity.Dish) error {
	return s.st.AddDish(dish)
}

func (s *DishService) UpdateDish(id int, dish *entity.Dish) error {
	return s.st.UpdateDish(id, dish)
}

func (s *DishService) DeleteDish(id int) error {
	return s.st.DeleteDish(id)
}
