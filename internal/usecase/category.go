package usecase

import (
	"cookbook/internal/storage"
	_ "github.com/lib/pq"
)

type CategoryService struct {
	st storage.Category
}

func NewCategoryService(st storage.Category) *CategoryService {
	return &CategoryService{st:st}
}