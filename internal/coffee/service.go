package coffee

import (
	"gogetters/internal/models"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCoffee(coffee *models.Coffee) error {

	if coffee.Late == "" {
		return errors.New(" Cpu cannot be empty")
	}

	if len(coffee.Late) < 2 {
		return errors.New("cannot cpu name must be at least 2 characters")
	}
	if coffee.Amount == 0 {
		return errors.New(" Cpu cannot be empty")
	}
	return s.repo.CreateCoffee(coffee)
}

func (s *Service) GetAllCoffee() ([]models.Coffee, error) {
	return s.repo.GetAllCoffee()
}

func (s *Service) UpdateCoffee(id uint, coffee *models.Coffee) error {

	if coffee.Late == "" {
		return errors.New(" Cpu cannot be empty")
	}

	if len(coffee.Late) < 2 {
		return errors.New("cannot cpu name must be at least 2 characters")
	}
	if coffee.Amount == 0 {
		return errors.New(" Cpu cannot be empty")
	}
	
	return s.repo.UpdateCoffee(id, coffee)
}


func (s *Service) DeleteCoffe(id uint) error {
	if id == 0 {
		return errors.New("laptop id must be greater than zero")
	}
	return s.repo.DeleteCoffe(id)
}
