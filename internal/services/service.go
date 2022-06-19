package services

import "github.com/XXena/shorter/internal/entities"

type Service struct {
	Record
}
type Record interface {
	Create(entities.Record) (shortURL string, err error)
	GetByURL(longURL string) (shortURL string, err error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}
