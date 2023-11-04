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

func (s *DishService) AddDish(dish *entity.Dish) (int, error) {
	return s.st.AddDish(dish)
}

func (s *DishService) UpdateDish(id int, dish *entity.Dish) error {
	return s.st.UpdateDish(id, dish)
}

func (s *DishService) DeleteDish(id int) error {
	return s.st.DeleteDish(id)
}

func (s *DishService) GetDishCousine(id int) ([]entity.Dish, error) {
	return s.st.GetDishCousine(id)
}

func (s *DishService) GetDishCategory(id int) ([]entity.Dish, error) {
	return s.st.GetDishCategory(id)
}

func (s *DishService) GetDishCousineCategory(cousineID, categoryID int) ([]entity.Dish, error) {
	return s.st.GetDishCousineCategory(cousineID, categoryID)
}