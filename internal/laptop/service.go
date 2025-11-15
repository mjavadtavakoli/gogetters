package laptop

import (
	"gogetters/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateLaptop(laptop *models.Laptop) error {
	return s.repo.CreateLaptop(laptop)
}

func (s *Service) GetAllLaptop() ([]models.Laptop, error) {
	return s.repo.GetAllLaptop()
}