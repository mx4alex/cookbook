package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
	"context"
	"unicode"
	"strings"
)

type DishService struct {
	st storage.Dish
}

func NewDishService(st storage.Dish) *DishService {
	return &DishService{st: st}
}

func (s *DishService) GetAllDishes(ctx context.Context) ([]entity.Dish, error) {
	return s.st.GetAllDishes(ctx)
}

func (s *DishService) GetDishInfo(ctx context.Context, id int) (entity.Dish, error) {
	return s.st.GetDishInfo(ctx, id)
}

func (s *DishService) AddDish(ctx context.Context, dish *entity.Dish) (int, error) {
	return s.st.AddDish(ctx, dish)
}

func (s *DishService) UpdateDish(ctx context.Context, id int, dish *entity.Dish) error {
	return s.st.UpdateDish(ctx, id, dish)
}

func (s *DishService) DeleteDish(ctx context.Context, id int) error {
	return s.st.DeleteDish(ctx, id)
}

func (s *DishService) GetDishCousine(ctx context.Context, id int) ([]entity.Dish, error) {
	return s.st.GetDishCousine(ctx, id)
}

func (s *DishService) GetDishCategory(ctx context.Context, id int) ([]entity.Dish, error) {
	return s.st.GetDishCategory(ctx, id)
}

func (s *DishService) GetDishCousineCategory(ctx context.Context, cousineID, categoryID int) ([]entity.Dish, error) {
	return s.st.GetDishCousineCategory(ctx, cousineID, categoryID)
}

func (s *DishService) GetDishSearch(ctx context.Context, text string) ([]entity.Dish, error) {
	var words []string
    var word string
	
	text = strings.ToLower(text)

    for _, r := range text {
        if unicode.IsSpace(r) || r == ',' {
            if len(word) > 0 {
                words = append(words, word)
                word = ""
            }

        } else {
			if len(word) == 0 {
				r = unicode.ToTitle(r)
			}
            word += string(r)
        }
    }
	
	if len(word) > 0 {
		words = append(words, word)
	}

	return s.st.GetDishSearch(ctx, words)
}