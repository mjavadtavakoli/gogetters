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

func (r *Repository) UpdateLaptop(id uint, laptop *models.Laptop) error {
	return r.DB.Model(&models.Laptop{}).Where("id = ?", id).Updates(laptop).Error
}


func (r *Repository) DeleteLaptop(id uint) error {
	return r.DB.Delete(&models.Laptop{}, id).Error
}

func (r *Repository) FindByCpu(cpu string) (*models.Laptop, error) {
	var laptop models.Laptop
	err := r.DB.Where("cpu = ?", cpu).First(&laptop).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &laptop, nil
}
