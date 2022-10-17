package handlers

import (
	"net/http"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/XXena/shorter/internal/services"
)

type handler struct {
	service services.RecordServiceInterface
	logger  logger.Interface
}
type HandlerInterface interface {
	Fetch(w http.ResponseWriter, r *http.Request)
	Redirect(w http.ResponseWriter, r *http.Request)
	InitRoutes() *http.ServeMux
}

func NewHandler(s services.RecordServiceInterface, l logger.Interface) HandlerInterface {
	return &handler{
		service: s,
		logger:  l,
	}
}

func (h *handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Redirect)
	mux.HandleFunc("/send", h.Fetch)
	return mux
}
