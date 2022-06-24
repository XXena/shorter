package handlers

import (
	"net/http"

	"github.com/XXena/shorter/internal/services"
)

type Handler struct {
	service *services.Service
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	//router := httprouter.New()
	//router.GET("/", h.fetch)
	mux := http.NewServeMux()
	//Service := services.NewService()
	//Handler := NewHandler(Service)
	mux.HandleFunc("/", h.Fetch)
	return mux
}
