package coffee

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

func (r *Repository) CreateCoffee(coffee *models.Coffee) error {
	return r.DB.Create(coffee).Error
}

func (r *Repository) GetAllCoffee() ([]models.Coffee, error) {
	var coffees []models.Coffee
	err := r.DB.Find(&coffees).Error
	return coffees, err
}

func (r *Repository) UpdateCoffee(id uint, coffee *models.Coffee) error {
	return r.DB.Model(&models.Coffee{}).Where("id = ?", id).Updates(coffee).Error
}
