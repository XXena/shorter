package handlers

import (
	"net/http"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/XXena/shorter/internal/services"
)

type Handler struct {
	service *services.Service
	logger  logger.Interface
}

func NewHandler(s *services.Service, l logger.Interface) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Redirect)
	mux.HandleFunc("/send", h.Fetch)
	return mux
}

//type Service struct {
//	RecordRepo
//	logger logger.Interface
//}
//type RecordRepo interface {
//	Create(entities.Record) (string, error)
//	GetByURL(string) (string, error)
//	Redirect(string) (string, error)
//	Update(recordID int, record entities.Record) error
//	Delete(recordID int) error
//}
//
//func NewService(r *repository.Repository, l logger.Interface) *Service {
//	return &Service{
//		RecordRepo: NewRecordService(r.Record, l),
//	}
//}
