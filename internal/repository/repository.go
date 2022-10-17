package repository

import (
	"github.com/XXena/shorter/internal/entities"
)

//type Repository struct {
//	RecordInterface
//}

type RecordInterface interface {
	Create(record entities.Record) (id int, err error)
	GetByURL(longURL string) (record entities.Record, err error)
	GetByToken(token string) (record entities.Record, err error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}

//
//func NewRepository(db *pgx.Conn, l logger.Interface) RecordInterface {
//	return &recordPostgres{
//		RecordInterface: NewRecordPostgres(db, l),
//	}
//}
