package services

import (
	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/internal/repository"
	"github.com/XXena/shorter/pkg/logger"
)

type Service struct {
	Record
	logger logger.Interface
}
type Record interface {
	Create(entities.Record) (string, error)
	GetByURL(string) (string, error)
	Redirect(string) (string, error)
	Update(recordID int, record entities.Record) error
	Delete(recordID int) error
}

func NewService(r *repository.Repository, l logger.Interface) *Service {
	return &Service{
		Record: NewRecordService(r.Record, l),
	}
}
