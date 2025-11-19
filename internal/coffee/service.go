package coffee

import (
	"gogetters/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCoffee(coffee *models.Coffee) error {
	return s.repo.CreateCoffee(coffee)
}

func (s *Service) GetAllCoffee() ([]models.Coffee, error) {
	return s.repo.GetAllCoffee()
}

func (s *Service) UpdateCoffee(id uint, coffee *models.Coffee) error {
	return s.repo.UpdateCoffee(id, coffee)
}