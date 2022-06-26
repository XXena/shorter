package repository

import (
	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/pkg/logger"
	"github.com/jackc/pgx"
)

type Repository struct {
	Record
}

type Record interface {
	Create(record entities.Record) (id int, err error)
	GetByURL(longURL string) (record entities.Record, err error)
	GetByToken(token string) (record entities.Record, err error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}

func NewRepository(db *pgx.Conn, l logger.Interface) *Repository {
	return &Repository{
		Record: NewRecordPostgres(db, l),
	}
}
