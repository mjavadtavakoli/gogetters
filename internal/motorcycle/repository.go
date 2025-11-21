package motorcycle

import (
	"errors"
	"gogetters/internal/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strings"
)

// ErrDuplicateKey is returned when a duplicate key constraint violation occurs
var ErrDuplicateKey = errors.New("duplicate key")

// RepositoryInterface defines the repository methods
type RepositoryInterface interface {
	CreateMotorcycle(motorcycle *models.Motorcycle) error
	GetAllMotorcycle() ([]models.Motorcycle, error)
	UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error
	DeleteMotorcycle(id uint) error
	FindByBrand(brand string) (*models.Motorcycle, error)
	FindByFueltype(fueltype string) (*models.Motorcycle, error)
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateMotorcycle(motorcycle *models.Motorcycle) error {
	err := r.DB.Create(motorcycle).Error
	if err != nil {
		// Check if error is a duplicate key violation (PostgreSQL error code 23505)
		// GORM might wrap the error, so we need to check both directly and unwrap
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrDuplicateKey
		}
		// Also check error string as fallback
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return ErrDuplicateKey
		}
		return err
	}
	return nil
}

func (r *Repository) GetAllMotorcycle() ([]models.Motorcycle, error) {
	var motorcycles []models.Motorcycle
	err := r.DB.Find(&motorcycles).Error
	return motorcycles, err
}

func (r *Repository) UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error {
	err := r.DB.Model(&models.Motorcycle{}).Where("id = ?", id).Updates(motorcycle).Error
	if err != nil {
		// Check if error is a duplicate key violation (PostgreSQL error code 23505)
		// GORM might wrap the error, so we need to check both directly and unwrap
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrDuplicateKey
		}
		// Also check error string as fallback
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return ErrDuplicateKey
		}
		return err
	}
	return nil
}

func (r *Repository) DeleteMotorcycle(id uint) error {
	return r.DB.Delete(&models.Motorcycle{}, id).Error
}

func (r *Repository) FindByBrand(brand string) (*models.Motorcycle, error) {
	var motorcycle models.Motorcycle
	err := r.DB.Where("brand = ?", brand).First(&motorcycle).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &motorcycle, nil
}

func (r *Repository) FindByFueltype(fueltype string) (*models.Motorcycle, error) {
	var motorcycle models.Motorcycle
	err := r.DB.Where("fueltype = ?", fueltype).First(&motorcycle).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &motorcycle, nil
}


