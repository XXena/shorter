package services

import (
	"time"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/internal/repository"
	"github.com/XXena/shorter/pkg/logger"
)

type Service struct {
	RecordRepo
	logger logger.Interface
}
type RecordRepo interface {
	Create(entities.Record) (string, error)
	ForwardToCreate(url string, expiry time.Time) (id []byte, err error)
	GetByURL(string) (string, error)
	Redirect(string) (string, error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}

func NewService(r *repository.Repository, l logger.Interface) *Service {
	return &Service{
		RecordRepo: NewRecordService(r.Record, l),
	}
}
