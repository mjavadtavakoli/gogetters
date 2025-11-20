package motorcycle

import (
	"errors"
	"gogetters/internal/models"
)

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMotorcycle(motorcycle *models.Motorcycle) error {

	if motorcycle.Price < 10000 {
		return errors.New("motorcycle price must be at least 10000")
	}

	if motorcycle.Totalspeed > 90 {
		return errors.New("cant use select speed after 90")
	}

	if len(motorcycle.Brand) < 3 {
		return errors.New("cannot brand name must be at least 3 characters")
	}

	if motorcycle.Brand == "" {
		return errors.New("motorcycle brand cannot be empty")
	}

	if motorcycle.Totalspeed == 0 {
		return errors.New("motorcycle totalspeed cannot be empty")
	}

	if motorcycle.Fueltype == "" {
		return errors.New("motorcycle fueltype cannot be empty")
	}

	if motorcycle.Price == 0 {
		return errors.New("motorcycle price cannot be empty")
	}

	return s.repo.CreateMotorcycle(motorcycle)
}

func (s *Service) GetAllMotorcycle() ([]models.Motorcycle, error) {
	return s.repo.GetAllMotorcycle()
}

func (s *Service) UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error {
	if motorcycle.Price < 100 {
		return errors.New("motorcycle price must be at least 10000")
	}
	if motorcycle.Price < 10000 {
		return errors.New("motorcycle price must be at least 10000")
	}

	if motorcycle.Totalspeed > 90 {
		return errors.New("cant use select speed after 90")
	}

	if len(motorcycle.Brand) < 3 {
		return errors.New("cannot brand name must be at least 3 characters")
	}

	if motorcycle.Brand == "" {
		return errors.New("motorcycle brand cannot be empty")
	}

	if motorcycle.Totalspeed == 0 {
		return errors.New("motorcycle totalspeed cannot be empty")
	}

	if motorcycle.Fueltype == "" {
		return errors.New("motorcycle fueltype cannot be empty")
	}

	if motorcycle.Price == 0 {
		return errors.New("motorcycle price cannot be empty")
	}

	return s.repo.UpdateMotorcycle(id, motorcycle)
}

func (s *Service) DeleteMotorcycle(id uint) error {
	if id == 0 {
		return errors.New("motorcycle id must be greater than zero")
	}
	return s.repo.DeleteMotorcycle(id)
}
