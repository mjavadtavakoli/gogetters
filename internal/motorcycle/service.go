package motorcycle

import (
	"gogetters/internal/models"
)

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMotorcycle(motorcycle *models.Motorcycle) error {
	return s.repo.CreateMotorcycle(motorcycle)
}

func (s *Service) GetAllMotorcycle() ([]models.Motorcycle, error) {
	return s.repo.GetAllMotorcycle()
}

func (s *Service) UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error {
	return s.repo.UpdateMotorcycle(id, motorcycle)
}