package repository

import (
	"github.com/XXena/shorter/internal/entities"
)

type RecordInterface interface {
	Create(record entities.Record) (id int, err error)
	GetByURL(longURL string) (record entities.Record, err error)
	GetByToken(token string) (record entities.Record, err error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}
