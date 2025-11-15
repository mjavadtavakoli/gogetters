package motorcycle

import (
	"gogetters/internal/models"
	"gorm.io/gorm"
)

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
