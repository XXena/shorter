package services

import (
	"time"

	"github.com/XXena/shorter/internal/entities"
)

type RecordServiceInterface interface {
	Create(entities.Record) (string, error)
	ForwardToCreate(url string, expiry time.Time) (id []byte, err error)
	GetByURL(string) (string, error)
	Redirect(string) (string, error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}
