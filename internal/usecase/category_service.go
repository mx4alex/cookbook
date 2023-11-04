package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
	_ "github.com/lib/pq"
	"context"
)

type CategoryService struct {
	st storage.Category
}

func NewCategoryService(st storage.Category) *CategoryService {
	return &CategoryService{st:st}
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]entity.Category, error) {
	return s.st.GetCategories(ctx)
}

func (s *CategoryService) AddCategory(ctx context.Context, category *entity.Category) (int, error) {
	return s.st.AddCategory(ctx, category)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id int, category *entity.Category) error {
	return s.st.UpdateCategory(ctx, id, category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) error {
	return s.st.DeleteCategory(ctx, id)
}