package service

import (
	"github.com/DenHax/LibreLib-back/internal/domain/models"
	"github.com/DenHax/LibreLib-back/internal/repo"
	"github.com/DenHax/LibreLib-back/internal/service/book"
)

type Book interface {
	Book(id int) (models.Book, error) // GET
	Create(song models.Book) (int, error)
	Delete(id int)
}

type Service struct {
	Book
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Book: book.NewBookService(repos.Book),
	}
}
