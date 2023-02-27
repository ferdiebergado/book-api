package repository

import (
	"book-api/entity"
)

type BookRepository struct {
	Repository
}

func (repo *BookRepository) Save(book *entity.Book) error {

	return nil

}
