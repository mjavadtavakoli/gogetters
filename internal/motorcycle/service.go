package motorcycle

import (
	"gogetters/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMotorcycle(motorcycle *models.Motorcycle) error {
	return s.repo.CreateMotorcycle(motorcycle)
}

func (s *Service) GetAllMotorcycle() ([]models.Motorcycle, error) {
	return s.repo.GetAllMotorcycle()
}