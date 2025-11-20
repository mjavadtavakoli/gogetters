package motorcycle

import (
	"gogetters/internal/models"
	"gorm.io/gorm"
)

// RepositoryInterface defines the repository methods
type RepositoryInterface interface {
	CreateMotorcycle(motorcycle *models.Motorcycle) error
	GetAllMotorcycle() ([]models.Motorcycle, error)
	UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error
	DeleteMotorcycle(id uint) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateMotorcycle(motorcycle *models.Motorcycle) error {
	return r.DB.Create(motorcycle).Error
}

func (r *Repository) GetAllMotorcycle() ([]models.Motorcycle, error) {
	var motorcycles []models.Motorcycle
	err := r.DB.Find(&motorcycles).Error
	return motorcycles, err
}

func (r *Repository) UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error {
	return r.DB.Model(&models.Motorcycle{}).Where("id = ?", id).Updates(motorcycle).Error
}

func (r *Repository) DeleteMotorcycle(id uint) error {
	return r.DB.Delete(&models.Motorcycle{}, id).Error
}
