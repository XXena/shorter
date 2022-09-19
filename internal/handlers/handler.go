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
type HandlerInterface interface {
	Fetch(w http.ResponseWriter, r *http.Request)
	Redirect(w http.ResponseWriter, r *http.Request)
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
