package laptop

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

func (r *Repository) CreateLaptop(laptop *models.Laptop) error {
	return r.DB.Create(laptop).Error
}

func (r *Repository) GetAllLaptop() ([]models.Laptop, error) {
	var laptops []models.Laptop
	err := r.DB.Find(&laptops).Error
	return laptops, err
}
