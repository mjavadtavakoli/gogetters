package book

import "gogetters/internal/models"

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

