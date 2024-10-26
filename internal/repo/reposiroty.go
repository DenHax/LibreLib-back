package repo

import (
	"github.com/DenHax/LibreLib-back/internal/domain/models"
	"github.com/DenHax/LibreLib-back/internal/storage/postgres"

	"github.com/DenHax/LibreLib-back/internal/repo/postgres/book"
)

type Book interface {
	Book(id int) (models.Book, error) // GET
	Create(song models.Book) (int, error)
	Delete(id int) error
	GetAll() ([]models.Book, error)
}

type Repository struct {
	Book
}

func NewRepository(s *postgres.Storage) *Repository {
	return &Repository{
		Book: book.NewBookPostgres(s),
	}
}
