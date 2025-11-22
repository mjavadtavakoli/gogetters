package book

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

func (r *Repository) Create(book *models.Book) error {
    return r.DB.Create(book).Error
}

func (r *Repository) GetAll() ([]models.Book, error) {
    var books []models.Book
    err := r.DB.Find(&books).Error
    return books, err
}


func (r *Repository) UpdateBook(id uint, book *models.Book) error {
	return r.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
}


func (r *Repository) DeleteBook(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}