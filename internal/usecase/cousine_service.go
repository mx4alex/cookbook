package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
	_ "github.com/lib/pq"
	"context"
)

type CousineService struct {
	st storage.Cousine
}

func NewCousineService(st storage.Cousine) *CousineService {
	return &CousineService{st:st}
}

func (s *CousineService) GetCousines(ctx context.Context) ([]entity.Cousine, error) {
	return s.st.GetCousines(ctx)
}

func (s *CousineService) AddCousine(ctx context.Context, cousine *entity.Cousine) (int, error) {
	return s.st.AddCousine(ctx, cousine)
}

func (s *CousineService) UpdateCousine(ctx context.Context, id int, cousine *entity.Cousine) error {
	return s.st.UpdateCousine(ctx, id, cousine)
}

func (s *CousineService) DeleteCousine(ctx context.Context, id int) error {
	return s.st.DeleteCousine(ctx, id)
}