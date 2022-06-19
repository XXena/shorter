package repository

import "github.com/XXena/shorter/internal/entities"

type Repository struct {
	Record
}

type Record interface {
	Create(recordID int, record entities.Record) (shortURL string, err error)
	//GetById(recordID int) (entities.Record, error)
	GetByURL(longURL string) (shortURL string, err error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}
