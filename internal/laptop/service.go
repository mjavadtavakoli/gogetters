package laptop

import (
	"errors"
	"gogetters/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateLaptop(laptop *models.Laptop) error {

	if laptop.Cpu == "" {
		return errors.New(" Cpu cannot be empty")
	}

	if len(laptop.Cpu) < 2 {
		return errors.New("cannot cpu name must be at least 2 characters")
	}


	return s.repo.CreateLaptop(laptop)
}

func (s *Service) GetAllLaptop() ([]models.Laptop, error) {
	return s.repo.GetAllLaptop()
}

func (s *Service) UpdateLaptop(id uint, laptop *models.Laptop) error {

	if laptop.Cpu == "" {
		return errors.New(" Cpu cannot be empty")
	}

	if len(laptop.Cpu) < 2 {
		return errors.New("cannot cpu name must be at least 2 characters")
	}
	return s.repo.UpdateLaptop(id, laptop)
}

func (s *Service) DeleteLaptop(id uint) error {
	if id == 0 {
		return errors.New("laptop id must be greater than zero")
	}
	return s.repo.DeleteLaptop(id)
}
