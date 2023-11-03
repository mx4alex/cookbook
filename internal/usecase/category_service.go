package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
	_ "github.com/lib/pq"
)

type CategoryService struct {
	st storage.Category
}

func NewCategoryService(st storage.Category) *CategoryService {
	return &CategoryService{st:st}
}

func (s *CategoryService) GetCategories() ([]entity.Category, error) {
	return s.st.GetCategories()
}

func (s *CategoryService) AddCategory(category *entity.Category) (int, error) {
	return s.st.AddCategory(category)
}

func (s *CategoryService) UpdateCategory(id int, category *entity.Category) error {
	return s.st.UpdateCategory(id, category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.st.DeleteCategory(id)
}