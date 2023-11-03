package usecase

import (
	"cookbook/internal/entity"
	"cookbook/internal/storage"
	_ "github.com/lib/pq"
)

type CousineService struct {
	st storage.Cousine
}

func NewCousineService(st storage.Cousine) *CousineService {
	return &CousineService{st:st}
}

func (s *CousineService) GetCousines() ([]entity.Cousine, error) {
	return s.st.GetCousines()
}

func (s *CousineService) AddCousine(cousine *entity.Cousine) (int, error) {
	return s.st.AddCousine(cousine)
}

func (s *CousineService) UpdateCousine(id int, cousine *entity.Cousine) error {
	return s.st.UpdateCousine(id, cousine)
}

func (s *CousineService) DeleteCousine(id int) error {
	return s.st.DeleteCousine(id)
}