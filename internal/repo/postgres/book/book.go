package book

import (
	"fmt"

	"github.com/DenHax/LibreLib-back/internal/domain/models"
	"github.com/DenHax/LibreLib-back/internal/storage/postgres"
	_ "github.com/lib/pq"
)

type BookPostgres struct {
	storage *postgres.Storage
}

func NewBookPostgres(s *postgres.Storage) *BookPostgres {
	return &BookPostgres{storage: s}
}

func (r *BookPostgres) Book(id int) (models.Book, error) {
	var song models.Book
	query := fmt.Sprintf(`SELECT id, song, group, lyrics FROM %s`,
		postgres.BookTable)
	if err := r.storage.DB.Get(&song, query); err != nil {
		return song, err
	}
	return song, nil
}

func (r *BookPostgres) GetAll() ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf(`SELECT id, song, group, lyrics FROM %s`,
		postgres.BookTable)
	if err := r.storage.DB.Select(&books, query); err != nil {
		return nil, err
	}
	return books, nil

}

func (r *BookPostgres) Create(book models.Book) (int, error) {
	tx, err := r.storage.DB.Begin()
	if err != nil {
		return 0, err
	}
	var songId int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (title) values ($1) RETURNING id", postgres.BookTable)

	row := tx.QueryRow(createBookQuery, book.Name)
	err = row.Scan(&songId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return songId, tx.Commit()
}

func (r *BookPostgres) Delete(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`,
		postgres.BookTable)
	_, err := r.storage.DB.Exec(query, id)
	return err
}
