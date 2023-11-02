package usecase

import (
	"cookbook/internal/storage"
	_ "github.com/lib/pq"
)

type CousineService struct {
	st storage.Cousine
}

func NewCousineService(st storage.Cousine) *CousineService {
	return &CousineService{st:st}
}