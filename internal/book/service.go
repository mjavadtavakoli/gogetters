package book

import (
	"gogetters/internal/models"
	"errors"
)

type Service struct {
    repo *Repository
}

func NewService(r *Repository) *Service {
    return &Service{repo: r}
}

func (s *Service) CreateBook(book *models.Book) error {
    return s.repo.Create(book)

    
}

func (s *Service) ListBooks() ([]models.Book, error) {
    return s.repo.GetAll()
}


func (s *Service) UpdateBook(id uint, book *models.Book) error {
	return s.repo.UpdateBook(id, book)
}


func (s *Service) DeleteBook(id uint) error {
	if id == 0 {
		return errors.New("book id must be greater than zero")
	}
	return s.repo.DeleteBook(id)
}
