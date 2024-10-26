package book

import (
	"github.com/DenHax/LibreLib-back/internal/domain/models"
	"github.com/DenHax/LibreLib-back/internal/repo"
)

type BookService struct {
	repo repo.Book
}

func NewBookService(repo repo.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Book(id int) (models.Book, error) {
	return models.Book{}, nil
}
func (s *BookService) Create(book models.Book) (int, error) {
	return 0, nil
}
func (s *BookService) Delete(id int) {

}
